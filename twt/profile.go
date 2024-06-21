package twt

import (
	"context"
	"encoding/json"
	"time"

	"freefrom.space/nobot/biz/dal/sqlite"
	"freefrom.space/nobot/core"
	"freefrom.space/nobot/core/twitterprofile"
	"freefrom.space/nobot/nostr_fixture"
	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

type RelayStatus struct {
	RelayURL string
	Status   string
}

type Metadata struct {
	Name        string `json:"name,omitempty"`
	About       string `json:"about,omitempty"`
	Picture     string `json:"picture,omitempty"`
	NIP05       string `json:"nip05,omitempty"`
	Banner      string `json:"banner,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	LUD16       string `json:"lud16,omitempty"`
	UserName    string `json:"username,omitempty"`
	Website     string `json:"website,omitempty"`
	Location    string `json:"location,omitempty"`
	Src         string `json:"src,omitempty"`
	IsBot       bool   `json:"is_bot,omitempty"`
}

var UpsertProfileCmd = &cli.Command{
	Name:    "profile-upsert",
	Aliases: []string{"pus"},
	Usage:   "从推特获取profile到本地并保存",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "username",
			Aliases: []string{"u"},
			// Required: true,
			Usage: "twitter user name, 可以重复设置,一次获取多位用户的profile",
		},
	},
	Action: func(cliCtx *cli.Context) error {
		// open db
		dbClient := sqlite.Client
		defer func(client *core.Client) {
			_ = client.Close()
		}(dbClient)

		// login twitter api
		scraper, err := LoginWithPredefined()
		if err != nil {
			return err
		}

		// fetch profiles and save to db
		userNameList := cliCtx.StringSlice("username")
		for _, userName := range userNameList {
			if err := FetchProfileToLocal(dbClient, scraper, userName); err != nil {
				log.Err(err).
					Str("username", userName).
					Msg("upsert profile failed")
			}
			time.Sleep(time.Second * 1)
		}
		return nil
	},
}

var ProfileUpdateCmd = &cli.Command{
	Name:    "profile-update",
	Aliases: []string{"pu"},
	Usage:   "对已经获取到本地的profile 创建 nostr key, 或更新profile event到relay",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "cmd",
			Aliases: []string{"c"},
			Usage: `
create-nostr: create keys
update-profile: update current profile and publish to relays
`,
		},
		RelayFlag,
	},
	Action: func(cliCtx *cli.Context) error {
		// open db
		relayList := cliCtx.StringSlice("relay")

		dbClient := sqlite.Client
		defer func(client *core.Client) {
			_ = client.Close()
		}(dbClient)
		ctx := context.Background()

		cmd := cliCtx.String("cmd")
		switch cmd {
		default:
			log.Info().Str("cmd", cmd).Msg("unknown cmd")
		case "update-profile":

			log.Info().Msg("create nostr profile")

			pq := dbClient.TwitterProfile.Query().
				Where(twitterprofile.NsecKeyNotNil(),
					twitterprofile.PublishNostrEQ(true),
				)

			profileList, err := pq.
				All(ctx)
			if err != nil {
				return err
			}
			for _, profile := range profileList {
				log.Info().
					Str("display name", profile.Name).
					Str("user name", profile.UserName).
					Msg("create nostr profile")

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
					Tags:      []nostr.Tag{[]string{}},
					Content:   string(metadataJsonStr),
				}
				_ = metaEvent.Sign(profile.PrivateKeyHex)

				relayPool := nostr.NewSimplePool(context.Background())
				status, err := nostr_fixture.PublishSignedEventToRelay(context.Background(), relayPool, relayList, metaEvent)
				if err != nil {
					log.Err(err).Msg("update profile failed")
				} else {
					log.Info().
						Bool("publish at lease one", status).
						Str("event", metaEvent.String()).
						Msg("update profile to relay ok")
				}
				nostr_fixture.CloseRelayPool(relayPool)
			}
		case "create-nostr":
			log.Info().Msg("create nostr keys for twitter profiles")

			pq := dbClient.TwitterProfile.Query().
				Where(twitterprofile.NsecKeyIsNil())

			profileList, err := pq.
				All(ctx)
			if err != nil {
				return err
			}
			for _, profile := range profileList {
				log.Info().
					Str("display name", profile.Name).
					Str("user name", profile.UserName).
					Msg("profile nostr keys missing")

				secKeyHex := nostr.GeneratePrivateKey()
				secKeyNIP19, _ := nip19.EncodePrivateKey(secKeyHex)
				pubKeyHex, _ := nostr.GetPublicKey(secKeyHex)
				pubKeyNIP19, _ := nip19.EncodePublicKey(pubKeyHex)
				err = dbClient.TwitterProfile.UpdateOne(profile).
					SetNsecKey(secKeyNIP19).
					SetNpubKey(pubKeyNIP19).
					SetPrivateKeyHex(secKeyHex).
					SetPublicKeyHex(pubKeyHex).
					SetPublishNostr(false).
					Exec(ctx)
				if err != nil {
					log.Err(err).Msg("update profile failed")
				}
			}
		}

		return nil
	},
}

// ProfileListCmd represents the command for listing local profiles.
var ProfileListCmd = &cli.Command{
	Name:    "profile-list",
	Aliases: []string{"pl"},
	Usage:   "list local profiles",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "show all profiles, default only show active profiles",
		},
	},
	Action: func(cliCtx *cli.Context) error {
		// open db
		dbClient := sqlite.Client
		defer func(client *core.Client) {
			_ = client.Close()
		}(dbClient)
		ctx := context.Background()

		showAll := cliCtx.Bool("all")

		pq := dbClient.TwitterProfile.Query()

		if !showAll {
			pq.Where(twitterprofile.ActiveEQ(true))
		}

		profileList, err := pq.
			All(ctx)
		if err != nil {
			return err
		}
		for _, profile := range profileList {
			log.Info().
				Str("display name", profile.Name).
				Str("user name", profile.UserName).
				Int("profile note count", profile.TwitterCount).
				Int("followers count", profile.FollowersCount).
				Bool("active (是否启用推文自动更新)?", profile.Active).
				Msg("profile")
		}

		return nil
	},
}
