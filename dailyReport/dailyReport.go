package dailyReport

import (
	"database/sql"
	"fmt"
	"regexp"
	"time"

	"freefrom.space/nobot/pkl"

	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

var conf = pkl.GetConf()

// CreateDailyReport retrieves public keys and usernames from the configuration,
// establishes a connection to the database, and creates an Excel file containing
// data fetched by executing SQL queries. The Excel file is saved with the name
// "result.xlsx". The function also applies pairwise conditional formatting to
// highlight specific cells based on certain conditions.
func CreateDailyReport() error {
	publicKeys := conf.NewConfig.GetReport().GetPublicKeys()
	userNames := conf.NewConfig.GetReport().GetUsernames()

	// Database connection
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Create an Excel file
	f := excelize.NewFile()
	defer f.SaveAs("result.xlsx")

	for i, publicKey := range publicKeys {
		userName := userNames[i]

		// SQL Query
		query := fmt.Sprintf(`
        WITH RankedNostr AS (
          SELECT
            publish_at,
            context,
            ROW_NUMBER() OVER (ORDER BY publish_at DESC) AS rn
          FROM
            nostr_notes_fetch_histories
          WHERE
            public_key_hex = '%s'
        ),
        RankedTwitter AS (
          SELECT
            text,
            strftime('%%Y-%%m-%%d %%H:%%M:%%S', post_at, '+9 hours') AS post_at_converted,
            ROW_NUMBER() OVER (ORDER BY post_at DESC) AS rn
          FROM
            twitter_notes
          WHERE
            user_name = '%s'
        )
        SELECT
          nn.publish_at,
          nn.context,
          tn.post_at_converted,
          tn.text
        FROM
          RankedNostr nn
        INNER JOIN
          RankedTwitter tn ON nn.rn = tn.rn;`, publicKey, userName)

		// Execute the query
		rows, err := db.Query(query)
		if err != nil {
			return err
		}
		defer rows.Close()

		sheetName := userName
		f.NewSheet(sheetName)
		setHeaders(f, sheetName)
		// the front left corner of the wall is the origin, and the coordinates of the back right are (800,700)
		// we below consider a particle moving in the region R. the size is particle is zero
		// the absolute value of the velocity of a particle does not change.
		// suppose that a particle is moving with velocity  (vx, vy)

		var freeformTexts, twitterTexts, publishAts, postAtConverteds []string

		rowIndex := 2
		for rows.Next() {
			var publishAt, freeform, postAtConverted, twitter string
			err = rows.Scan(&publishAt, &freeform, &postAtConverted, &twitter)
			if err != nil {
				return err
			}

			freeformTexts = append(freeformTexts, freeform)
			twitterTexts = append(twitterTexts, twitter)
			publishAts = append(publishAts, publishAt)
			postAtConverteds = append(postAtConverteds, postAtConverted)

			setCellValues(f, sheetName, rowIndex, publishAt, freeform, postAtConverted, twitter)
			rowIndex++
		}

		applyPairwiseConditionalFormatting(f, sheetName, freeformTexts, twitterTexts, publishAts, postAtConverteds)
	}

	if err := f.SaveAs("result.xlsx"); err != nil {
		return err
	}
	log.Info().Msg("Excel file with multiple sheets successfully created")
	return nil
}

// setHeaders sets the headers for the specified sheet in an Excel file.
// It takes a pointer to an `excelize.File` object and the name of the sheet as parameters.
func setHeaders(f *excelize.File, sheetName string) {
	headers := []string{"Publish At", "Freeform", "Post At Converted", "Twitter"}
	for colIndex, header := range headers {
		cell := fmt.Sprintf("%c%d", 'A'+colIndex, 1)
		f.SetCellValue(sheetName, cell, header)
	}
}

// setCellValues sets the cell values in the specified sheet of an Excel file.
// The cell values are determined by the provided parameters.
//
// Parameters:
// - f: An excelize.File representing the Excel file.
// - sheetName: A string indicating the name of the sheet.
// - rowIndex: An integer representing the row index of the cell.
// - publishAt: A string representing the Publish At value.
// - freeform: A string representing the Freeform value.
// - postAtConverted: A string representing the Post At Converted value.
// - twitter: A string representing the Twitter value.
//
// This function sets the cell values in the specified Excel sheet using the provided
// parameters. It uses the excelize.SetCellValue() function to set the values in the
// cells with the format "<columnLetter><rowIndex>".
//
// Example Usage:
//
//	setCellValues(f, "Sheet1", 2, "2022-01-01", "Some text", "2022-01-01 12:34:56", "Tweet content")
//
// Note: This function is used within the CreateDailyReport() function to populate the
// cells of an Excel file with fetched data.
func setCellValues(f *excelize.File, sheetName string, rowIndex int, publishAt, freeform, postAtConverted, twitter string) {
	f.SetCellValue(sheetName, fmt.Sprintf("A%d", rowIndex), publishAt)
	f.SetCellValue(sheetName, fmt.Sprintf("B%d", rowIndex), freeform)
	f.SetCellValue(sheetName, fmt.Sprintf("C%d", rowIndex), postAtConverted)
	f.SetCellValue(sheetName, fmt.Sprintf("D%d", rowIndex), twitter)
}

func applyPairwiseConditionalFormatting(f *excelize.File, sheetName string, freeformTexts, twitterTexts, publishAt, postAtConverted []string) {
	blueStyle, err := f.NewStyle(`{"font":{"color":"#0000FF"}}`)
	checkErr(err)
	redStyle, err := f.NewStyle(`{"font":{"color":"#FF0000"}}`)
	checkErr(err)
	buleTodayCount := 0
	redTodayCount := 0
	for i, freeform := range freeformTexts {
		if !hasSimilar(freeform, twitterTexts, 0.80) {
			f.SetCellStyle(sheetName, fmt.Sprintf("B%d", i+2), fmt.Sprintf("B%d", i+2), blueStyle)
			if isTodayRFC3339(publishAt[i]) {
				buleTodayCount++
			}
		}
	}

	for i, twitter := range twitterTexts {
		if !hasSimilar(twitter, freeformTexts, 0.80) {
			f.SetCellStyle(sheetName, fmt.Sprintf("D%d", i+2), fmt.Sprintf("D%d", i+2), redStyle)
			if isToday(postAtConverted[i]) {
				redTodayCount++
			}

		}
	}
	log.Info().Msgf("blueTodayCount:", buleTodayCount)
	log.Info().Msgf("redTodayCount:", redTodayCount)
	f.SetCellValue(sheetName, "E1", "blueTodayCount")
	f.SetCellValue(sheetName, "E2", buleTodayCount)
	f.SetCellValue(sheetName, "F1", "redTodayCount")
	f.SetCellValue(sheetName, "F2", redTodayCount)
}

func hasSimilar(text string, texts []string, threshold float64) bool {
	textClean := removeURLs(text)
	for _, t := range texts {
		tClean := removeURLs(t)
		if isSimilar(textClean, tClean, threshold) {
			return true
		}
	}
	return false
}

func removeURLs(text string) string {
	re := regexp.MustCompile(`http[s]?://\S+`)
	return re.ReplaceAllString(text, "")
}

func isSimilar(a, b string, threshold float64) bool {
	distance := levenshtein.DistanceForStrings([]rune(a), []rune(b), levenshtein.DefaultOptions)
	longerLength := max(len(a), len(b))
	if longerLength == 0 {
		return true // both strings are empty
	}
	similarity := 1 - float64(distance)/float64(longerLength)
	return similarity >= threshold
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func checkErr(err error) {
	if err != nil {
		return
	}
}

func isToday(dateStr string) bool {
	t, err := time.Parse("2006-01-02 15:04:05", dateStr)
	checkErr(err)

	now := time.Now()
	return t.Year() == now.Year() && t.Month() == now.Month() && t.Day() == now.Day()
}

func isTodayRFC3339(dateStr string) bool {
	t, err := time.Parse(time.RFC3339, dateStr)
	checkErr(err)

	now := time.Now()
	return t.Year() == now.Year() && t.Month() == now.Month() && t.Day() == now.Day()
}
