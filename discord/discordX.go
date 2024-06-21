package discordX

import (
	"os"

	"freefrom.space/nobot/conf"
	pkg "freefrom.space/nobot/video"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func UploadVideo(s *discordgo.Session, videoPath string) string {
	// convert the video format to .mp4
	convertedVideoPath := "./convertedVideo.mp4"
	pkg.ConvertVideoFormat(videoPath, convertedVideoPath)

	// open the converted video file
	file, err := os.Open(convertedVideoPath)
	if err != nil {
		log.Err(err).Msgf("error opening video file,", err)
	}
	defer file.Close()

	// create a new file message
	message := &discordgo.MessageSend{
		Files: []*discordgo.File{
			{
				Name:   convertedVideoPath,
				Reader: file,
			},
		},
	}

	// send the message with the video file
	msg, err := s.ChannelMessageSendComplex(conf.GetConf().Discord.ChannelId, message)
	if err != nil {
		log.Err(err).Msgf("error sending message,", err)
	}
	// return the URL of the video
	return msg.Attachments[0].URL
}

func DownloadAndUploadVideoToDiscord(videoURL string) (string, error) {
	pkg.DownloadVideo(videoURL)
	dg, err := discordgo.New("Bot " + conf.GetConf().Discord.Token)
	if err != nil {
		log.Err(err).Msgf("error creating Discord session,", err)
		return "", err
	}

	videoLink := UploadVideo(dg, "downloaded_video.mp4")
	log.Info().Msgf("Video link:", videoLink)
	return videoLink, nil
}
