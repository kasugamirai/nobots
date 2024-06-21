package pkg

import (
	"io"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
)

func DownloadVideo(url string) {
	// sent GET request to video URL
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// create a file to hold the video
	out, err := os.Create("downloaded_video.mp4")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// copy the response body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}

	log.Info().Msg("video downloaded")
}
