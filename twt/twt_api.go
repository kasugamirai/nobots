package twt

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"freefrom.space/nobot/biz/dal/postgres"
	"freefrom.space/nobot/core"
	"freefrom.space/nobot/core/twitternote"
	"freefrom.space/nobot/core/twitterprofile"
	"freefrom.space/nobot/nostr_fixture"
	"github.com/carlmjohnson/requests"
	"github.com/go-faster/errors"
	twitterscraper "github.com/n0madic/twitter-scraper"
	"github.com/nbd-wtf/go-nostr"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

const (
	authHeaderKey = "authorization"
	auth          = "Bearer AAAAAAAAAAAAAAAAAAAAAFQODgEAAAAAVHTp76lzh3rFzcHbmHVvQxYYpTw%3DckAlMINMjmCwxUcaXbAN4XqJVdgMJaHqNOFgPMK0zN1qLqLQCF"

	apiRoot = "https://api.twitter.com"

	graphql = apiRoot + "/graphql/"

	graphUser                 = graphql + "u7wQyGi6oExe8_TRWGMq4Q/UserResultByScreenNameQuery"
	graphUserById             = graphql + "oPppcargziU1uDQHAUmH-A/UserResultByIdQuery"
	graphUserTweets           = graphql + "3JNH4e9dq1BifLxAa3UMWg/UserWithProfileTweetsQueryV2"
	graphUserTweetsAndReplies = graphql + "8IS8MaO-2EN6GZZZb8jF0g/UserWithProfileTweetsAndRepliesQueryV2"
	graphUserMedia            = graphql + "PDfFf8hGeJvUCiTyWtw4wQ/MediaTimelineV2"
	graphTweet                = graphql + "83h5UyHZ9wEKBVzALX8R_g/ConversationTimelineV2"
	graphTweetResult          = graphql + "sITyJdhRPpvpEjg4waUmTA/TweetResultByIdQuery"
	graphSearchTimeline       = graphql + "gkjsKepM6gl_HmFWoWKfgg/SearchTimeline"
	graphListById             = graphql + "iTpgCtbdxrsJfyx0cFjHqg/ListByRestId"
	graphListBySlug           = graphql + "-kmqNvm5Y-cVrfvBy6docg/ListBySlug"
	graphListMembers          = graphql + "P4NpVZDqUD_7MEM84L-8nw/ListMembers"
	graphListTweets           = graphql + "BbGLL1ZfMibdFNWlk7a0Pw/ListTimeline"
	charset                   = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

const (
	features = `{
"android_graphql_skip_api_media_color_palette": false,
"blue_business_profile_image_shape_enabled": false,
"creator_subscriptions_subscription_count_enabled": false,
"creator_subscriptions_tweet_preview_api_enabled": true,
"freedom_of_speech_not_reach_fetch_enabled": false,
"graphql_is_translatable_rweb_tweet_is_translatable_enabled": false,
"hidden_profile_likes_enabled": false,
"highlights_tweets_tab_ui_enabled": false,
"interactive_text_enabled": false,
"longform_notetweets_consumption_enabled": true,
"longform_notetweets_inline_media_enabled": false,
"longform_notetweets_richtext_consumption_enabled": true,
"longform_notetweets_rich_text_read_enabled": false,
"responsive_web_edit_tweet_api_enabled": false,
"responsive_web_enhance_cards_enabled": false,
"responsive_web_graphql_exclude_directive_enabled": true,
"responsive_web_graphql_skip_user_profile_image_extensions_enabled": false,
"responsive_web_graphql_timeline_navigation_enabled": false,
"responsive_web_media_download_video_enabled": false,
"responsive_web_text_conversations_enabled": false,
"responsive_web_twitter_article_tweet_consumption_enabled": false,
"responsive_web_twitter_blue_verified_badge_is_enabled": true,
"rweb_lists_timeline_redesign_enabled": true,
"spaces_2022_h2_clipping": true,
"spaces_2022_h2_spaces_communities": true,
"standardized_nudges_misinfo": false,
"subscriptions_verification_info_enabled": true,
"subscriptions_verification_info_reason_enabled": true,
"subscriptions_verification_info_verified_since_enabled": true,
"super_follow_badge_privacy_enabled": false,
"super_follow_exclusive_tweet_notifications_enabled": false,
"super_follow_tweet_api_enabled": false,
"super_follow_user_api_enabled": false,
"tweet_awards_web_tipping_enabled": false,
"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled": false,
"tweetypie_unmention_optimization_enabled": false,
"unified_cards_ad_metadata_container_dynamic_card_content_query_enabled": false,
"verified_phone_label_enabled": false,
"vibe_api_enabled": false,
"view_counts_everywhere_api_enabled": false
}`
)

// TwtUserName sends a request to the Twitter API to fetch user information by screen name.
// The function uses a predefined set of headers and parameters for the request.
// It logs the response received from the API using log.Info.
// The function returns an error if the request fails.
func TwtUserName() error {
	ctx := context.Background()

	myClient := *http.DefaultClient
	myClient.Jar = requests.NewCookieJar()

	var resp string
	err := requests.URL(graphUser).
		Client(&myClient).
		Transport(requests.Record(nil, "rr")).
		Header(authHeaderKey, auth).
		Header("connection", "keep-alive").
		Header("content-type", "application/json").
		Header("x-twitter-active-user", "yes").
		Header("authority", "api.twitter.com").
		// Header("accept-encoding", "gzip").
		Header("accept", "*/*").
		Header("DNT", "1").
		Header("accept-language", "en-US,en;q=0.9").
		Param("variables", `{ "rest_id": "cloudwu", "screen_name":"cloudwu", "count": 30 }`).
		// Param("variables", `{ "rest_id": "davidxifeng", "screen_name":"davidxifeng", "count": 30 }`).
		Param("features", features).
		CheckStatus(http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusBadRequest).
		ToString(&resp).
		Fetch(ctx)
	if err != nil {
		log.Err(err).Msg("fetch err")
	}

	// log.Info().Str("resp", resp).Msg("ok")
	log.Info().Msg(resp)
	return nil
}

// twitterAccount represents a Twitter account with the following properties:
// - Username: The username of the account
// - Password: The password of the account
// - Email: The email associated with the account
// - Code: The verification code for two-factor authentication
// LoginAndSaveCookies is a method that allows the twitterAccount to login to Twitter and save the cookies.
// It uses the provided username, password, and email to login to Twitter using a twitterscraper.
// If the login is successful, it retrieves the cookies and saves them to a file named "cookies1.json".
// If the login fails or the cookies cannot be retrieved, an error is returned.
// LoginAndSaveCookiesTwoFactor is similar to LoginAndSaveCookies, but it uses the verification code for two-factor authentication instead of the email.
// Example usage:
//
//	account := &twitterAccount{
//	    Username: "example",
//	    Password: "password",
//	    Email:    "example@example.com",
//	}
//
// err := account.LoginAndSaveCookies()
//
//	if err != nil {
//	    fmt.Println("Failed to login and save cookies:", err)
//	}
//
// Example usage with two-factor authentication:
//
//	account := &twitterAccount{
//	    Username: "example",
//	    Password: "password",
//	    Code:     "123456",
//	}
//
// err := account.LoginAndSaveCookiesTwoFactor()
//
//	if err != nil {
//	    fmt.Println("Failed to login and save cookies with two-factor authentication:", err)
//	}
type twitterAccount struct {
	Username string
	Password string
	Email    string
	Code     string
}

// LoginAndSaveCookies logs in to Twitter using the provided username, password, and email.
// It creates a new scraper object and calls the Login method with the provided credentials.
// If the login is successful, it retrieves the cookies from the scraper using the GetCookies method.
// If the cookies are not nil, it marshals them into JSON format and writes them to a file named "cookies1.json".
// If any error occurs during the login process or writing the cookies to file, an error is returned.
// Example usage:
//
//	err := account.LoginAndSaveCookies()
func (twt *twitterAccount) LoginAndSaveCookies() error {
	scraper := twitterscraper.New()
	err := scraper.Login(twt.Username, twt.Password, twt.Email)
	if err != nil {
		return err
	}

	cookies := scraper.GetCookies()
	if cookies != nil {
		js, _ := json.Marshal(cookies)
		f, _ := os.Create("cookies1.json")
		_, err = f.Write(js)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Failed to get cookies")
	}

	return nil
}

// LoginAndSaveCookiesTwoFactor logs in to Twitter using two-factor authentication.
// It creates a new scraper object and calls the Login method with the provided username, password, and code.
// If the login is successful, it retrieves the cookies from the scraper using the GetCookies method.
// If the cookies are not nil, it marshals them into JSON format and writes them to a file named "cookies1.json".
// If any error occurs during the login process or writing the cookies to file, an error is returned.
// Example usage:
//
//	err := account.LoginAndSaveCookiesTwoFactor()
func (twt *twitterAccount) LoginAndSaveCookiesTwoFactor() error {
	scraper := twitterscraper.New()
	err := scraper.Login(twt.Username, twt.Password, twt.Code)
	if err != nil {
		return err
	}
	cookies := scraper.GetCookies()
	if cookies != nil {
		js, _ := json.Marshal(cookies)
		f, _ := os.Create("cookies1.json")
		_, err = f.Write(js)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Failed to get cookies")
	}
	return nil
}

// LoginWithPredefined function logs in to Twitter using a predefined cookie.
// If the cookie directory exists and contains cookie files, a random cookie file is selected.
// If the selected cookie file exists, the function loads the cookies from the file and sets them in the scraper.
// If the selected cookie file does not exist, the function performs a new login and saves the cookies to the file.
// The function returns a scraper object that is logged in, or an error if the login fails.
func LoginWithPredefined() (*twitterscraper.Scraper, error) {
	cookieDir := "conf/twt_users_cookies"
	var selectedFile string
	files, err := os.ReadDir(cookieDir)
	if err != nil {
		log.Info().Msg("cookies Dir not found, login & save cookie")
	} else {
		var cookieFiles []string
		for _, file := range files {
			cookieFiles = append(cookieFiles, file.Name())
		}
		if len(cookieFiles) == 0 {
			log.Err(err).Msg("no cookies found")
		}
		// Seed the random number generator
		src := rand.NewSource(time.Now().UnixNano())
		r := rand.New(src)
		selectedFile = cookieFiles[r.Intn(len(cookieFiles))]
	}
	log.Info().Msg(selectedFile)

	cookieFile := "cookies1.json"
	scraper := twitterscraper.New()
	//scraper.UpdateTransport(func(transport http.RoundTripper) http.RoundTripper {
	//	return requests.Record(transport, "twt_rr")
	//})

	f, err := os.Open("conf/twt_users_cookies/" + selectedFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Err(err).Msg("cookies file not found, login & save cookie")

			username := os.Getenv("TWITTER_USERNAME")
			password := os.Getenv("TWITTER_PASSWORD")
			email := os.Getenv("TWITTER_EMAIL")
			err := scraper.Login(username, password, email)
			if err != nil {
				return nil, err
			}
			log.Info().Bool("logged in ok", scraper.IsLoggedIn()).Msg("login ok")
			cookies := scraper.GetCookies()
			if cookies != nil {
				js, _ := json.Marshal(cookies)
				f, _ := os.Create(cookieFile)
				_, err = f.Write(js)
				if err != nil {
					log.Warn().
						Err(err).
						Msg("write cookies json failed")
					return nil, err
				}
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		var cookies []*http.Cookie
		_ = json.NewDecoder(f).Decode(&cookies)
		scraper.SetCookies(cookies)
	}
	if !scraper.IsLoggedIn() {
		return nil, err
	}
	log.Info().Msg("login ok")
	scraper.WithDelay(3)
	// scraper.WithReplies(true)
	return scraper, nil
}

func FetchProfileToLocal(dbClient *core.Client,
	scraper *twitterscraper.Scraper, twtUserName string,
) error {
	profile, err := scraper.GetProfile(twtUserName)
	if err != nil {
		return err
	}
	log.Info().
		Str("name", profile.Name).
		Str("url", profile.URL).
		Interface("p", profile).
		Msg("profile")

	rawYamlStr, err := yaml.Marshal(profile)
	if err != nil {
		return err
	}

	ctx := context.Background()

	profileCreator := dbClient.TwitterProfile.Create().
		SetRawYaml(string(rawYamlStr)).
		SetUserName(profile.Username).
		SetName(profile.Name).
		SetUserID(profile.UserID).
		SetAvatar(profile.Avatar).
		SetBanner(profile.Banner).
		SetBiography(profile.Biography).
		SetBirthday(profile.Birthday).
		SetFollowersCount(profile.FollowersCount).
		SetFollowingCount(profile.FollowingCount).
		SetFriendsCount(profile.FriendsCount).
		SetListedCount(profile.ListedCount).
		SetLikesCount(profile.LikesCount).
		SetTwitterCount(profile.TweetsCount).
		SetNillableJoinedTime(profile.Joined).
		SetLocation(profile.Location).
		SetURL(profile.URL).
		SetWebsite(profile.Website)

	if len(profile.PinnedTweetIDs) > 0 {
		profileCreator.SetPinnedNoteID(profile.PinnedTweetIDs[0])
	}

	dbProfileID, err := profileCreator.
		OnConflictColumns(twitterprofile.FieldUserID).
		UpdateNewValues().
		ID(ctx)
	// upsert方法由于没有读db, 其结果是不确定的,所以没有返回完整对象的save方法. (create有)
	if err != nil {
		return err
	}

	log.Info().
		Str("username", twtUserName).
		Int("db id", dbProfileID).
		Msg("save to db ok")

	_ = dbClient.TwitterFetchHistory.Create().
		SetFetchType("profile").
		SetFetchAt(time.Now()).
		SetContent(string(rawYamlStr)).
		SetItemID(profile.UserID).
		SetName(profile.Name).
		Exec(ctx)

	return nil
}

// FetchTwitterToLocal fetches tweets from the Twitter API using a scraper and saves them to a local database.
// The function takes the following parameters:
// - dbClient: The database client to interface with the local database
// - scraper: The Twitter scraper used to fetch tweets
// - twtUserName: The username of the Twitter account to fetch tweets from
// - maxTweets: The maximum number of tweets to fetch
// The function starts by initializing the context and creating a counter for the fetched tweets.
// It also stores the current time as the fetched time.
// It then iterates over each tweet received from the scraper.
// If an error occurs while fetching the tweets, the function logs the error and returns it.
// Otherwise, it marshals the tweet into YAML format and saves it to the rawYamlStr variable.
// It also appends the ID of the fetched tweet to the thisFetchedNoteIDList.
// The function logs information about the fetched tweet using log.Info, including the tweet ID, text, name, count, and various flags.
// It then creates a TwitterNote struct using the dbClient and sets its properties using the fetched tweet data.
// This includes setting the ID, fetchedAt, rawYaml, userID, username, text, conversationID, permanentURL, isPin, isReply, isRetweet, retweetedStatusID, postAt, replyTo, sensitiveContent
func FetchTwitterToLocal(dbClient *core.Client,
	scraper *twitterscraper.Scraper, twtUserName string, maxTweets int,
) error {
	ctx := context.Background()
	i := 0

	fetchedTime := time.Now()

	var thisFetchedNoteIDList []string

	for note := range scraper.GetTweets(ctx, twtUserName, maxTweets) {
		if note.Error != nil {
			log.Err(note.Error).Msg("get tweets failed")
			return note.Error
		}

		rawYamlStr, err := yaml.Marshal(note)
		if err != nil {
			return err
		}

		thisFetchedNoteIDList = append(thisFetchedNoteIDList, note.ID)

		i += 1
		log.Info().
			Str("id", note.ID).
			Str("text", note.Text).
			Str("name", note.Name).
			Int("count", i).
			Bool("has retweeted ?", note.RetweetedStatus != nil).
			Bool("has quoted ?", note.QuotedStatus != nil).
			Bool("has in_reply_to ?", note.InReplyToStatus != nil).
			Int("thread len", len(note.Thread)).
			Time("post at", note.TimeParsed).
			Msg("[new twitter]")

		noteCreator := dbClient.TwitterNote.Create().
			SetID(note.ID).
			SetFetchedAt(fetchedTime).
			SetRawYaml(string(rawYamlStr)).
			SetUserID(note.UserID).
			SetUserName(note.Username).
			SetText(note.Text).
			SetConversationID(note.ConversationID).
			SetPermanentURL(note.PermanentURL).
			SetIsPin(note.IsPin).
			SetIsReply(note.IsReply).
			SetIsRetweet(note.IsRetweet).
			SetRetweetedStatusID(note.RetweetedStatusID).
			SetPostAt(note.TimeParsed).
			SetReplyTo(note.InReplyToStatusID).
			SetSensitiveContent(note.SensitiveContent).
			SetLikes(note.Likes).
			SetViews(note.Views).
			SetRetweets(note.Retweets).
			SetReplies(note.Replies).
			SetUrls(note.URLs).
			SetHTMLText(note.HTML).
			SetHashTags(note.Hashtags)

		var photoIDList []string
		var photoURLList []string
		for _, photo := range note.Photos {
			photoIDList = append(photoIDList, photo.ID)
			photoURLList = append(photoURLList, photo.URL)
		}
		noteCreator.SetPhotoIds(photoIDList).
			SetPhotoUrls(photoURLList)

		var videoIDList []string
		var videoPreviewList []string
		var videoURLList []string
		for _, video := range note.Videos {
			videoIDList = append(videoIDList, video.ID)
			videoURLList = append(videoURLList, video.URL)
			videoPreviewList = append(videoPreviewList, video.URL)
		}
		noteCreator.SetVideoID(videoIDList).
			SetVideoPreview(videoPreviewList).
			SetVideoUrls(videoURLList)

		var mentionID []string
		var mentionName []string
		var mentionUserName []string
		for _, mention := range note.Mentions {
			mentionID = append(mentionID, mention.ID)
			mentionName = append(mentionName, mention.Name)
			mentionUserName = append(mentionUserName, mention.Username)
		}
		noteCreator.SetMentionID(mentionID).
			SetMentionName(mentionName).
			SetMentionUsername(mentionUserName)

		err = noteCreator.OnConflictColumns(twitternote.FieldID).
			UpdateNewValues().
			Update(func(upsert *core.TwitterNoteUpsert) {
				// upsert.ClearFetchedAt()
				upsert.SetIgnore(twitternote.FieldFetchedAt)
			}).
			Exec(ctx)
		if err != nil {
			return err
		}

		minSleep := 300
		maxSleep := 1000
		sleepTime := rand.Intn(maxSleep-minSleep+1) + minSleep
		time.Sleep(time.Millisecond * time.Duration(sleepTime))
	}

	newCount, err := dbClient.TwitterNote.Query().
		Where(twitternote.IDIn(thisFetchedNoteIDList...),
			twitternote.FetchedAtEQ(fetchedTime)).
		Count(context.Background())
	if err != nil {
		return err
	}
	log.Info().
		Int("new note count", newCount).
		Msg("fetch status")

	return nil
}

func PublishProfile(profile *core.TwitterProfile,
	relayList []string,
) {
	metadataJsonStr, _ := json.Marshal(Metadata{
		Name:        profile.Name,
		DisplayName: profile.Name,
		UserName:    profile.UserName,
		Website:     profile.Website,
		Picture:     profile.Avatar,
		Banner:      profile.Banner,
		About:       profile.Biography,
		NIP05:       "",
		LUD16:       "",
		Location:    profile.Location,
		Src:         profile.URL,
	})

	metaEvent := &nostr.Event{
		CreatedAt: nostr.Now(),
		Kind:      nostr.KindProfileMetadata,
		Tags:      []nostr.Tag{},
		Content:   string(metadataJsonStr),
	}
	_ = metaEvent.Sign(profile.PrivateKeyHex)

	relayPool := nostr.NewSimplePool(context.Background())
	status, err := nostr_fixture.PublishSignedEventToRelay(context.Background(), relayPool, relayList, metaEvent)
	if err != nil {
		log.Err(err).Msg("update profile failed")
	} else {
		log.Info().
			Bool("publish at lease one ?", status).
			Str("event", metaEvent.String()).
			Msg("update profile to relay ok")
	}
	nostr_fixture.CloseRelayPool(relayPool)
}

func multiplySyncTwitterNostr(dbClient *core.Client, scraper *twitterscraper.Scraper,
	relayList []string, maxCount int, profiles []*core.TwitterProfile) {

	//create a pool for each profile
	var queues [][]*core.TwitterNote

	// for each profile, fetch notes and put them into a queue
	for _, profile := range profiles {
		if profile != nil {
			_, err := FetchTwitterNotes(dbClient, scraper, profile, maxCount)
			if err != nil {
				log.Err(err).Msg("failed to fetch twitter notes")
			}
		} else {
			log.Warn().Msg("Profile is nil, skipping FetchTwitterNotes")
		}
		randSleep := rand.Intn(21) + 10
		time.Sleep(time.Duration(randSleep) * time.Second)
	}
	// one pool for one profile, or else there may be error from relay
	relayPool := nostr.NewSimplePool(context.Background())
	for len(queues) > 0 {
		for i := 0; i < len(queues); i++ {
			if len(queues[i]) > 0 {
				note := queues[i][0]
				queues[i] = queues[i][1:]
				textEvent := nostr_fixture.ConvertTwitterContentToNostrEvent(note, profiles[i])
				publishOK, err := nostr_fixture.PublishSignedEventToRelay(context.Background(),
					relayPool, relayList, textEvent)
				if err != nil {
					log.Err(err).Msg("publish text note event failed")
					continue
				}
				if publishOK {
					_ = dbClient.TwitterNote.UpdateOne(note).
						SetEventID(textEvent.ID).
						SetEvent(textEvent.String()).
						SetPublishedAt(textEvent.CreatedAt.Time()).
						Exec(context.Background())

					log.Info().
						Str("text", note.Text).
						Bool("publishOK", publishOK).
						Str("name", note.UserName).
						Str("event id", textEvent.ID).
						Msg("publish note to nostr ok")
				}
			}
		}
		// remove empty queues
		var nonEmptyQueues [][]*core.TwitterNote
		for _, queue := range queues {
			if len(queue) > 0 {
				nonEmptyQueues = append(nonEmptyQueues, queue)
			}
		}
		queues = nonEmptyQueues
	}
	nostr_fixture.CloseRelayPool(relayPool)
}

func SyncTwitterNostr(dbClient *core.Client, scraper *twitterscraper.Scraper,
	relayList []string, maxCount int, profile *core.TwitterProfile,
) {
	postgresClient := postgres.Client
	log.Info().
		Str("display name", profile.Name).
		Str("user name", profile.UserName).
		Int("profile note count", profile.TwitterCount).
		Int("followers count", profile.FollowersCount).
		Msg("publish to nostr")

	noteList, err := FetchTwitterNotes(dbClient, scraper, profile, maxCount)
	if err != nil {
		log.Err(err).Msg("fetch twitter note failed")
		return
	}
	// one pool for one profile, or else there may be error from relay
	relayPool := nostr.NewSimplePool(context.Background())
	for _, note := range noteList {
		textEvent := nostr_fixture.ConvertTwitterContentToNostrEvent(note, profile)
		publishOK, err := nostr_fixture.PublishSignedEventToRelay(context.Background(),
			relayPool, relayList, textEvent)
		if err != nil {
			log.Err(err).Msg("publish text note event failed")
			continue
		}
		if publishOK {
			// TODO 记录详细的发布结果
			_ = dbClient.TwitterNote.UpdateOne(note).
				SetEventID(textEvent.ID).
				SetEvent(textEvent.String()).
				SetPublishedAt(textEvent.CreatedAt.Time()).
				Exec(context.Background())

			// Seed the random number generator
			rand.Seed(time.Now().UnixNano())

			// Generate a random bigint
			randomBigInt := rand.Int63()

			err = postgresClient.Events.Create().
				SetPubkey([]byte("123456789")).
				SetEventID([]byte(textEvent.ID)).
				SetKind(1).
				SetContent("freefrom_test").
				SetSig([]byte(textEvent.String())).
				SetID(randomBigInt).
				SetCreatedAt(time.Now()).
				Exec(context.Background())
			if err != nil {
				log.Err(err).Msg("failed to save event to postgres")
			}

			log.Info().
				Str("text", note.Text).
				Bool("publishOK", publishOK).
				Str("name", note.UserName).
				Str("event id", textEvent.ID).
				Msg("publish note to nostr ok")
		}
		time.Sleep(5 * time.Minute)
	}
	nostr_fixture.CloseRelayPool(relayPool)
}

func FetchTwitterNotes(dbClient *core.Client, scraper *twitterscraper.Scraper,
	twtProfile *core.TwitterProfile, maxTweets int,
) ([]*core.TwitterNote, error) {
	ctx := context.Background()
	fetchedCount := 0

	fetchedTime := time.Now()
	var thisFetchedNoteIDList []string
	for note := range scraper.GetTweets(ctx, twtProfile.UserName, maxTweets) {
		if note.Error != nil {
			if strings.Contains(note.Error.Error(), "429 Too Many Requests") {
				log.Err(note.Error).Msg("Rate limit exceeded, sleeping for 30 minutes")
				time.Sleep(30 * time.Minute)
				continue // after sleeping, continue with the next iteration
			} else {
				log.Err(note.Error).Msg("get tweets failed")
				break
			}
		}
		rawYamlStr, err := yaml.Marshal(note)
		if err != nil {
			log.Err(err).Msg("note to yaml failed")
			rawYamlStr = make([]byte, 0)
		}

		thisFetchedNoteIDList = append(thisFetchedNoteIDList, note.ID)

		fetchedCount += 1
		log.Info().
			Str("id", note.ID).
			Str("text", note.Text).
			Str("name", note.Name).
			Int("count", fetchedCount).
			Time("post at", note.TimeParsed).
			Msg("[twitter note fetched]")

		noteCreator := dbClient.TwitterNote.Create().
			SetID(note.ID).
			SetFetchedAt(fetchedTime).
			SetRawYaml(string(rawYamlStr)).
			SetUserID(note.UserID).
			SetUserName(note.Username).
			SetText(note.Text).
			SetConversationID(note.ConversationID).
			SetPermanentURL(note.PermanentURL).
			SetIsPin(note.IsPin).
			SetIsReply(note.IsReply).
			SetIsRetweet(note.IsRetweet).
			SetRetweetedStatusID(note.RetweetedStatusID).
			SetPostAt(note.TimeParsed).
			SetReplyTo(note.InReplyToStatusID).
			SetSensitiveContent(note.SensitiveContent).
			SetLikes(note.Likes).
			SetViews(note.Views).
			SetRetweets(note.Retweets).
			SetReplies(note.Replies).
			SetUrls(note.URLs).
			SetHTMLText(note.HTML).
			SetHashTags(note.Hashtags)

		var photoIDList []string
		var photoURLList []string
		for _, photo := range note.Photos {
			photoIDList = append(photoIDList, photo.ID)
			photoURLList = append(photoURLList, photo.URL)
		}
		noteCreator.SetPhotoIds(photoIDList).
			SetPhotoUrls(photoURLList)

		videoIDList := make([]string, 0)
		videoURLList := make([]string, 0)
		videoPreviewList := make([]string, 0)

		for _, video := range note.Videos {
			videoIDList = append(videoIDList, video.ID)
			videoURLList = append(videoURLList, video.URL)
			videoPreviewList = append(videoPreviewList, video.URL)
		}
		// 暂时把 gif和video都存放到一起
		for _, gif := range note.GIFs {
			videoIDList = append(videoIDList, gif.ID)
			videoURLList = append(videoURLList, gif.URL)
			videoPreviewList = append(videoPreviewList, gif.URL)
		}
		noteCreator.SetVideoID(videoIDList).
			SetVideoPreview(videoPreviewList).
			SetVideoUrls(videoURLList)

		var mentionID []string
		var mentionName []string
		var mentionUserName []string
		for _, mention := range note.Mentions {
			mentionID = append(mentionID, mention.ID)
			mentionName = append(mentionName, mention.Name)
			mentionUserName = append(mentionUserName, mention.Username)
		}
		noteCreator.SetMentionID(mentionID).
			SetMentionName(mentionName).
			SetMentionUsername(mentionUserName)

		err = noteCreator.OnConflictColumns(twitternote.FieldID).
			UpdateNewValues().
			Update(func(upsert *core.TwitterNoteUpsert) {
				// upsert.ClearFetchedAt()
				upsert.SetIgnore(twitternote.FieldFetchedAt)
			}).
			Exec(ctx)
		if err != nil {
			break
		}

		minSleep := 300
		maxSleep := 1000
		sleepTime := rand.Intn(maxSleep-minSleep+1) + minSleep
		time.Sleep(time.Millisecond * time.Duration(sleepTime))
	}

	newNoteList, err := dbClient.TwitterNote.Query().
		Where(twitternote.IDIn(thisFetchedNoteIDList...),
			twitternote.FetchedAtEQ(fetchedTime)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	newCount := len(newNoteList)

	log.Info().
		Int("new note count", newCount).
		Msg("fetch notes done")

	err = dbClient.TwitterNotesFetchHistory.Create().
		SetFetchedAt(fetchedTime).
		SetUserID(twtProfile.UserID).
		SetUsername(twtProfile.UserName).
		SetName(twtProfile.Name).
		SetFetchCount(fetchedCount).
		SetNewCount(newCount).
		Exec(context.Background())
	if err != nil {
		log.Err(err).Msg("save notes fetch history failed")
	}
	return newNoteList, nil
}

func stringWithCharset(length int, charset string, r *rand.Rand) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	return string(b)
}

func String(length int, r *rand.Rand) string {
	return stringWithCharset(length, charset, r)
}

func CreateRandProfile() *core.TwitterProfile {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	secKeyHex := nostr.GeneratePrivateKey()
	pubKeyHex, _ := nostr.GetPublicKey(secKeyHex)

	return &core.TwitterProfile{
		Name:          String(10, r),                    // Random name of length 10
		UserName:      String(10, r),                    // Random username of length 10
		Website:       "https://apple.com",              // Example website
		Avatar:        "https://example.com/avatar.jpg", // Example avatar URL
		Banner:        "https://example.com/banner.jpg", // Example banner URL
		Biography:     String(50, r),                    // Random biography of length 50
		Location:      "Example Location",               // Example location
		URL:           "https://google.com",             // Example URL
		PrivateKeyHex: secKeyHex,
		PublicKeyHex:  pubKeyHex,
	}
}
