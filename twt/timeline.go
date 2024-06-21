package twt

import (
	"context"
	"encoding/csv"
	"github.com/nbd-wtf/go-nostr"
	"os"
	"os/signal"
	"syscall"
	"time"

	"freefrom.space/nobot/biz/dal/sqlite"
	"freefrom.space/nobot/dailyReport"
	"freefrom.space/nobot/nostr_fixture"
	"freefrom.space/nobot/pkl"

	"gopkg.in/yaml.v3"

	"github.com/go-co-op/gocron"
	twitterscraper "github.com/n0madic/twitter-scraper"
	"github.com/nbd-wtf/go-nostr/nip19"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"freefrom.space/nobot/core"
	"freefrom.space/nobot/core/twitternote"
	"freefrom.space/nobot/core/twitterprofile"
)

var conf = pkl.GetConf()

type VideoTransformRequest struct {
	VideoURL string `json:"video_url"`
}

type Response struct {
	Ok       bool   `json:"ok"`
	VideoURL string `json:"video_url"`
}

var TimeLineSimpleCmd = &cli.Command{
	Name:    "note-simple",
	Aliases: []string{"ns"},
	Usage:   "fetch notes by user id.",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:        "username",
			Aliases:     []string{"u"},
			Required:    false,
			DefaultText: "Chai20230817",
			Usage:       "twitter user name, 可以重复设置,一次获取多位用户的profile",
		},
		&cli.IntFlag{
			Name:    "count",
			Aliases: []string{"c"},
			Value:   5,
			Usage:   "max fetch count",
		},
	},
	Action: func(cliCtx *cli.Context) error {
		userNameList := cliCtx.StringSlice("username")
		maxCount := cliCtx.Int("count")
		verbose := cliCtx.Bool("verbose")

		// open db
		dbClient := sqlite.Client
		if verbose {
			dbClient = dbClient.Debug()
		}
		defer func(client *core.Client) {
			_ = client.Close()
		}(dbClient)

		// login twitter api
		scraper, err := LoginWithPredefined()
		if err != nil {
			return err
		}

		// fetch profiles and save it to db
		for _, userName := range userNameList {
			if err := FetchTwitterToLocal(dbClient, scraper, userName, maxCount); err != nil {
				log.Err(err).
					Str("username", userName).
					Msg("fetch twitter notes failed")
			}
			time.Sleep(time.Second * 1)
		}
		return nil
	},
}

var TimeLineCheckNewCmd = &cli.Command{
	Name:    "note-update",
	Aliases: []string{"nu"},
	Usage:   "fetch all active target twitter account top n note one by one.",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "count",
			Aliases: []string{"c"},
			Value:   3,
			Usage:   "max fetch count",
		},
		&cli.BoolFlag{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "load all known profile",
		},
	},
	Action: func(cliCtx *cli.Context) error {
		showAll := cliCtx.Bool("all")
		maxCount := cliCtx.Int("count")

		// open db
		dbClient := sqlite.Client
		defer func(client *core.Client) {
			_ = client.Close()
		}(dbClient)
		ctx := context.Background()

		pq := dbClient.TwitterProfile.Query()

		if !showAll {
			pq.Where(twitterprofile.ActiveEQ(true))
		}

		profileList, err := pq.
			All(ctx)
		if err != nil {
			return err
		}

		totalNoteCount, err := dbClient.TwitterNote.Query().Count(context.Background())
		if err != nil {
			return err
		}
		log.Info().
			Int("user count", len(profileList)).
			Int("check count", maxCount).
			Int("current note count", totalNoteCount).
			Msg("check new notes")

		scraper, err := LoginWithPredefined()
		if err != nil {
			return err
		}

		for _, profile := range profileList {
			// 更新profile?
			// 获取最新几条推文
			// 获取属于profile的新推文数量,并记录到获取历史中,这样还可以统计📉推主更新频率
			log.Info().
				Str("display name", profile.Name).
				Str("user name", profile.UserName).
				Int("profile note count", profile.TwitterCount).
				Int("followers count", profile.FollowersCount).
				Msg("check new tweets")
			if err := FetchTwitterToLocal(dbClient, scraper, profile.UserName, maxCount); err != nil {
				log.Err(err).
					Str("username", profile.UserName).
					Msg("fetch profile failed")
			}
			time.Sleep(time.Second * 2)
		}

		newTotalNoteCount, err := dbClient.TwitterNote.Query().Count(context.Background())
		if err != nil {
			return err
		}
		log.Info().
			Int("current note count", newTotalNoteCount).
			Int("new note count", newTotalNoteCount-totalNoteCount).
			Msg("fetch all profile done")

		return nil
	},
}

var SimplePublishCmd = &cli.Command{
	Name:    "note-publish",
	Aliases: []string{"np"},
	Usage:   "publish un published notes by username",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "count",
			Aliases: []string{"c"},
			Value:   0,
			Usage:   "limit max publish count, <= 0 means no limit",
		},
		RelayFlag,
		UserNameListFlag,
	},
	Action: func(cliCtx *cli.Context) error {
		maxCount := cliCtx.Int("count")
		relayList := cliCtx.StringSlice("relay")
		userNameList := cliCtx.StringSlice("username")

		// open db
		dbClient := sqlite.Client
		defer func(client *core.Client) {
			_ = client.Close()
		}(dbClient)
		ctx := context.Background()

		for _, username := range userNameList {
			profile, err := dbClient.TwitterProfile.Query().
				Where(twitterprofile.UserNameEQ(username)).
				Only(ctx)
			if err != nil {
				continue
			}

			qn := dbClient.TwitterNote.Query().
				Where(twitternote.UserIDEQ(profile.UserID),
					twitternote.EventIDIsNil())
			if maxCount > 0 {
				qn.Limit(maxCount)
			}
			noteList, err := qn.
				All(ctx)
			if err != nil {
				return err
			}

			// one pool for one profile, or else there may be error from relay
			relayPool := nostr.NewSimplePool(context.Background())
			for _, note := range noteList {
				textEvent := nostr_fixture.ConvertTwitterContentToNostrEvent(note, profile)
				publishOK, err := nostr_fixture.PublishSignedEventToRelay(context.Background(),
					relayPool, relayList, textEvent)
				if err != nil {
					continue
				}
				time.Sleep(time.Second * 3)
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
						Msg("publish note to nostr")
				}
			}

			nostr_fixture.CloseRelayPool(relayPool)
		}

		return nil
	},
}

var (
	RelayFlag = &cli.StringSliceFlag{
		Name:     "relay",
		Aliases:  []string{"r"},
		Required: false,
		Value:    cli.NewStringSlice("wss://freerelay.xyz"),
		Usage:    "target relay list, 可以重复设置. wss://freerelay.xyz",
	}

	VerboseFlag = &cli.BoolFlag{
		Name:    "verbose",
		Aliases: []string{"v"},
		Value:   false,
		Usage:   "show verbose log",
	}
	UserNameListFlag = &cli.StringSliceFlag{
		Name:     "username",
		Aliases:  []string{"u"},
		Required: false,
		Usage:    "twitter user name, 可以重复设置,一次处理多位用户的profile",
	}
)

var FetchPublishCmd = &cli.Command{
	Name:    "fetch-publish",
	Aliases: []string{"fp"},
	Usage: `fetch from twitter and publish to nostr, name and count required.
执行间隔在上一层任务编排时指定参数.
`,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "count",
			Aliases: []string{"c"},
			Value:   5,
			Usage:   "fetch and save publish max n count",
		},
		RelayFlag,
		UserNameListFlag,
		&cli.BoolFlag{
			Name:    "migrate",
			Aliases: []string{"m"},
			Value:   false,
			Usage:   "run db migrate on startup",
		},
	},
	Action: func(cliCtx *cli.Context) error {
		verbose := cliCtx.Bool("verbose")
		migrate := cliCtx.Bool("migrate")
		relayListFlag := cliCtx.StringSlice("relay")
		relayListFile := conf.NewConfig.GetRelays().GetUrls()
		relayList := append(relayListFlag, relayListFile...)

		userNameListFlag := cliCtx.StringSlice("username")
		userNameListFile := conf.NewConfig.GetTwitter().GetUsers()
		userNameList := append(userNameListFlag, userNameListFile...)
		maxCount := cliCtx.Int("count")

		// open db
		dbClient := sqlite.Client
		defer func(client *core.Client) {
			_ = client.Close()
		}(dbClient)
		if verbose {
			dbClient = dbClient.Debug()
		}
		if migrate {
			if err := dbClient.Schema.Create(context.Background()); err != nil {
				log.Err(err).Msg("migrate db failed")
				return cli.Exit(err, 1)
			} else {
				log.Info().Msg("migrate db ok")
			}
		}
		ctx := context.Background()

		// login twitter api
		scraper, err := LoginWithPredefined()
		if err != nil {
			return cli.Exit(err, 1)
		}
		for _, username := range userNameList {
			profile, err := dbClient.TwitterProfile.Query().
				Where(twitterprofile.UserNameEQ(username)).
				Only(ctx)
			if err == nil {
				SyncTwitterNostr(dbClient, scraper, relayList, maxCount, profile)
			} else if core.IsNotFound(err) {
				// try fetch profile & create nostr keys & publish profile to relays
				profile, err = ImportTwitterNostr(dbClient, scraper, relayList, username)
				if err == nil {
					SyncTwitterNostr(dbClient, scraper, relayList, maxCount, profile)
				}
			}
		}
		return nil
	},
}

var ShuffleFetchPublishCmd = &cli.Command{
	Name:    "shuffle-fetch-publish",
	Aliases: []string{"sfp"},
	Usage: `fetch from twitter and publish to nostr, name and count required.
执行间隔在上一层任务编排时指定参数.
`,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "count",
			Aliases: []string{"c"},
			Value:   5,
			Usage:   "fetch and save publish max n count",
		},
		RelayFlag,
		UserNameListFlag,
		&cli.BoolFlag{
			Name:    "migrate",
			Aliases: []string{"m"},
			Value:   false,
			Usage:   "run db migrate on startup",
		},
	},
	Action: func(cliCtx *cli.Context) error {
		verbose := cliCtx.Bool("verbose")
		migrate := cliCtx.Bool("migrate")
		relayListFlag := cliCtx.StringSlice("relay")
		relayListFile := conf.NewConfig.GetRelays().GetUrls()
		relayList := append(relayListFlag, relayListFile...)

		userNameListFlag := cliCtx.StringSlice("username")
		userNameListFile := conf.NewConfig.GetTwitter().GetUsers()
		userNameList := append(userNameListFlag, userNameListFile...)
		maxCount := cliCtx.Int("count")

		// open db
		dbClient := sqlite.Client
		defer func(client *core.Client) {
			_ = client.Close()
		}(dbClient)
		if verbose {
			dbClient = dbClient.Debug()
		}
		if migrate {
			if err := dbClient.Schema.Create(context.Background()); err != nil {
				log.Err(err).Msg("migrate db failed")
				return cli.Exit(err, 1)
			} else {
				log.Info().Msg("migrate db ok")
			}
		}
		ctx := context.Background()

		// login twitter api
		scraper, err := LoginWithPredefined()
		if err != nil {
			return cli.Exit(err, 1)
		}
		profiles := make([]*core.TwitterProfile, len(userNameList))
		for _, username := range userNameList {

			profile, err := dbClient.TwitterProfile.Query().
				Where(twitterprofile.UserNameEQ(username)).
				Only(ctx)

			if err == nil {
				profiles = append(profiles, profile)
			} else if core.IsNotFound(err) {
				// try fetch profile & create nostr keys & publish profile to relays
				profile, err = ImportTwitterNostr(dbClient, scraper, relayList, username)
				if err == nil {
					profiles = append(profiles, profile)
				}
			}
		}
		print("profiles", profiles)
		multiplySyncTwitterNostr(dbClient, scraper, relayList, maxCount, profiles)
		return nil
	},
}

func ImportTwitterNostr(dbClient *core.Client, scraper *twitterscraper.Scraper,
	relayList []string, username string,
) (*core.TwitterProfile, error) {
	profile, err := scraper.GetProfile(username)
	if err != nil {
		return nil, err
	}
	rawYamlStr, err := yaml.Marshal(profile)
	if err != nil {
		return nil, err
	}

	secKeyHex := nostr.GeneratePrivateKey()
	secKeyNIP19, _ := nip19.EncodePrivateKey(secKeyHex)
	pubKeyHex, _ := nostr.GetPublicKey(secKeyHex)
	pubKeyNIP19, _ := nip19.EncodePublicKey(pubKeyHex)

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
		// nostr keys
		SetNsecKey(secKeyNIP19).
		SetNpubKey(pubKeyNIP19).
		SetPrivateKeyHex(secKeyHex).
		SetPublicKeyHex(pubKeyHex).
		SetPublishNostr(true).
		SetWebsite(profile.Website)

	if len(profile.PinnedTweetIDs) > 0 {
		profileCreator.SetPinnedNoteID(profile.PinnedTweetIDs[0])
	}

	dbProfile, err := profileCreator.Save(ctx)
	if err != nil {
		return nil, err
	}

	_ = dbClient.TwitterFetchHistory.Create().
		SetFetchType("profile").
		SetFetchAt(time.Now()).
		SetContent(string(rawYamlStr)).
		SetItemID(profile.UserID).
		SetName(profile.Name).
		Exec(ctx)

	PublishProfile(dbProfile, relayList)

	return dbProfile, nil
}

var FinalTwtMirrorCmd = &cli.Command{
	Name:    "twt-mirror",
	Aliases: []string{"tm"},
	Usage:   `twitter mirror daemon`,
	Flags: []cli.Flag{
		RelayFlag,
		&cli.BoolFlag{
			Name:    "migrate",
			Aliases: []string{"m"},
			Value:   false,
			Usage:   "run db migrate on startup",
		},
	},
	Action: func(cliCtx *cli.Context) error {
		verbose := cliCtx.Bool("verbose")
		migrate := cliCtx.Bool("migrate")
		// relayList := cliCtx.StringSlice("relay")

		// open db
		dbClient := sqlite.Client
		defer func(client *core.Client) {
			_ = client.Close()
		}(dbClient)
		if verbose {
			dbClient = dbClient.Debug()
		}
		if migrate {
			if err := dbClient.Schema.Create(context.Background()); err != nil {
				log.Err(err).Msg("migrate db failed")
				return cli.Exit(err, 1)
			} else {
				log.Info().Msg("migrate db ok")
			}
		}

		// login Twitter api
		//scraper, err := LoginWithPredefined()
		//if err != nil {
		//	return cli.Exit(err, 1)
		//}
		//log.Info().Msg("login twitter api ok")
		tokyo, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			return cli.Exit(err, 1)
		}
		sched := gocron.NewScheduler(tokyo)
		sched.SingletonModeAll()
		log.Info().Msg("sched start")
		go sched.StartBlocking()
		job, err := sched.EveryRandom(3, 5).
			Seconds().
			Do(func() {
				log.Info().Msg("in go routine")
				time.Sleep(1)
			})
		if err != nil {
			log.Err(err).Msg("queue fl")
		}
		log.Info().Interface("job", job).Msg("job status")

		gracefulShutdown := make(chan os.Signal, 1)
		signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		log.Info().Msg("waiting for exit signal")
		<-gracefulShutdown

		log.Info().Msg("stop sched")
		sched.StopBlockingChan()

		log.Info().Msg("exiting main")
		// ctx, cancel := context.WithCancel(context.Background())
		// defer cancel()
		return nil
	},
}

var FetchAndSaveCmd = &cli.Command{
	Name:    "fetch-save",
	Aliases: []string{"fs"},
	Usage:   "fetch from twitter and save to database without publishing",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "count",
			Aliases: []string{"c"},
			Value:   5,
			Usage:   "fetch and save max count",
		},
		UserNameListFlag,
		&cli.BoolFlag{
			Name:    "migrate",
			Aliases: []string{"m"},
			Value:   false,
			Usage:   "run db migrate on startup",
		},
	},
	Action: func(cliCtx *cli.Context) error {
		migrate := cliCtx.Bool("migrate")
		userNameListFlag := cliCtx.StringSlice("username")
		userNameListFile := conf.NewConfig.GetTwitter().GetVipUsers()
		userNameList := append(userNameListFlag, userNameListFile...)
		maxCount := cliCtx.Int("count")

		// open db
		dbClient := sqlite.Client
		defer func(client *core.Client) {
			_ = client.Close()
		}(dbClient)

		if migrate {
			if err := dbClient.Schema.Create(context.Background()); err != nil {
				log.Err(err).Msg("migrate db failed")
				return cli.Exit(err, 1)
			} else {
				log.Info().Msg("migrate db ok")
			}
		}
		ctx := context.Background()

		// login twitter api
		scraper, err := LoginWithPredefined()
		if err != nil {
			log.Err(err).Msg("login twitter api failed")
			return cli.Exit(err, 1)
		}
		log.Info().Msg("login twitter api ok")

		// Iterate over each username and fetch their data
		for _, username := range userNameList {
			// Check if profile exists in the db, if not fetch and save
			profile, err := dbClient.TwitterProfile.Query().
				Where(twitterprofile.UserNameEQ(username)).
				Only(ctx)
			if err != nil && core.IsNotFound(err) {
				// Fetch and save the Twitter profile to the database
				profile, err = ImportTwitterNostr(dbClient, scraper, nil, username)
				if err != nil {
					log.Err(err).Msg("failed to fetch and save twitter profile")
					continue
				}
			}

			// Fetch and save notes (tweets) to the database
			if profile != nil {
				_, err := FetchTwitterNotes(dbClient, scraper, profile, maxCount)
				if err != nil {
					log.Err(err).Msg("failed to fetch twitter notes")
				}
			}
		}
		return nil
	},
}

var FetchFFNotesCmd = &cli.Command{
	Name:    "nostr-fetch",
	Aliases: []string{"nf"},
	Usage:   "start processing events and store them in the database",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:     "pubkey",
			Aliases:  []string{"pub"},
			Required: true,
			Usage:    "获取多位用户的nostr pub hex key",
		},
		RelayFlag,
	},

	Action: func(cliCtx *cli.Context) error {
		relayList := cliCtx.StringSlice("relay")
		pubKeyHex := cliCtx.StringSlice("pubkey")

		err := nostr_fixture.CatchUserDataFromRelay(relayList, pubKeyHex)
		if err != nil {
			log.Err(err).Msg("failed to catch user data from relay")
		}
		return nil
	},
}

var CreateDailyReportCmd = &cli.Command{
	Name:    "daily-report",
	Aliases: []string{"dr"},
	Usage:   "create daily report",
	Action: func(c *cli.Context) error {
		err := dailyReport.CreateDailyReport()
		if err != nil {
			log.Err(err).Msg("failed to create daily report")
		}
		return nil
	},
}

var CreateRandomAccountAndFollowCmd = &cli.Command{
	Name:    "create-Random-account-follow",
	Aliases: []string{"craf"},
	Usage:   "create a random account and follow someone",
	Action: func(c *cli.Context) error {
		// create nostr account
		secKeyHex := nostr.GeneratePrivateKey()
		pubKeyHex, _ := nostr.GetPublicKey(secKeyHex)

		log.Info().Msg("create nostr account ok")
		log.Info().Msg("privateKey:" + secKeyHex)
		log.Info().Msg("PublicKey:" + pubKeyHex)

		// create random profile
		profile := CreateRandProfile()

		// publish profile to nostr
		PublishProfile(profile, conf.NewConfig.GetRelays().GetUrls())

		// follow someone
		targetPubKey := conf.NewConfig.GetNostr().GetPublicKeys()
		followEvent := &nostr.Event{
			Kind:      nostr.KindContactList,
			PubKey:    pubKeyHex,
			Tags:      []nostr.Tag{{"p", targetPubKey}},
			CreatedAt: nostr.Now(),
		}

		_ = followEvent.Sign(secKeyHex)

		relayPool := nostr.NewSimplePool(context.Background())
		_, err := nostr_fixture.PublishSignedEventToRelay(context.Background(),
			relayPool, conf.NewConfig.GetRelays().GetUrls(), followEvent)
		if err != nil {
			log.Err(err).Msg("publish follow event failed")
		}
		log.Info().Msg("publish follow event ok")
		nostr_fixture.CloseRelayPool(relayPool)
		return nil
	},
}

var CreateAccountAndReactionCmd = &cli.Command{
	Name:    "create-account-reaction",
	Aliases: []string{"car"},
	Usage:   "create account and send reaction to someone",
	Action: func(c *cli.Context) error {
		// create nostr account
		secKeyHex := nostr.GeneratePrivateKey()
		pubKeyHex, _ := nostr.GetPublicKey(secKeyHex)

		log.
			Info().
			Msgf("privateKey:", secKeyHex)
		log.
			Info().
			Msgf("PublicKey:", pubKeyHex)

		// send message
		targetPubKey := conf.NewConfig.GetNostr().GetPublicKeys()
		targetEventID := conf.NewConfig.GetNostr().GetEventID()
		reactionEvent := &nostr.Event{
			Kind: nostr.KindReaction,
			Tags: []nostr.Tag{
				{"p", targetPubKey},
				{"e", targetEventID},
			},
			Content: "❤",
		}
		_ = reactionEvent.Sign(secKeyHex)

		relayPool := nostr.NewSimplePool(context.Background())
		_, err := nostr_fixture.PublishSignedEventToRelay(context.Background(),
			relayPool, conf.NewConfig.GetRelays().GetUrls(), reactionEvent)
		if err != nil {
			log.Err(err).Msg("publish reaction event failed")
		}
		log.Info().Msg("publish reaction event ok")
		nostr_fixture.CloseRelayPool(relayPool)
		return nil
	},
}

var GetUserFollowingCmd = &cli.Command{
	Name:    "get-user-following",
	Aliases: []string{"guf"},
	Usage:   "get user following list",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "publicKeyHex",
			Required: true,
			Aliases:  []string{"pub"},
			Usage:    "public key hex",
		},
	},

	Action: func(c *cli.Context) error {
		relayList := conf.NewConfig.GetRelays().GetUrls()
		publicKeyHex := c.String("publicKeyHex")
		ret, err := nostr_fixture.GetUserFollowings(relayList, publicKeyHex)
		if err != nil {
			log.Err(err).Msg("failed to get user following")
		}
		err = writeSliceToCSV("output.csv", ret)
		if err != nil {
			log.Err(err).Msg("failed to write to csv")
		}
		return nil
	},
}

var GetUserCookies = &cli.Command{
	Name:    "get-user-cookies",
	Aliases: []string{"guc"},
	Usage:   "get user cookies",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "Username",
			Required: true,
			Aliases:  []string{"u"},
			Usage:    "username",
		},
		&cli.StringFlag{
			Name:     "Password",
			Required: true,
			Aliases:  []string{"p"},
			Usage:    "password",
		},
		&cli.StringFlag{
			Name:     "Email",
			Required: false,
			Aliases:  []string{"e"},
			Usage:    "email",
		},
	},

	Action: func(c *cli.Context) error {
		// 创建一个 TwitterAccount 结构体实例
		account := &twitterAccount{
			Username: c.String("Username"),
			Password: c.String("Password"),
			Email:    c.String("Email"),
		}

		// 调用 LoginAndSaveCookies 方法
		err := account.LoginAndSaveCookies()
		if err != nil {
			return err
		}

		return nil
	},
}

var GetUserCookiesTwoFactor = &cli.Command{
	Name:    "get-user-cookies-two-factor",
	Aliases: []string{"guc2"},
	Usage:   "get user cookies two factor",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "Username",
			Required: true,
			Aliases:  []string{"u"},
			Usage:    "username",
		},
		&cli.StringFlag{
			Name:     "Password",
			Required: true,
			Aliases:  []string{"p"},
			Usage:    "password",
		},
		&cli.StringFlag{
			Name:     "Code",
			Required: false,
			Aliases:  []string{"c"},
			Usage:    "code",
		},
	},

	Action: func(c *cli.Context) error {
		// 创建一个 TwitterAccount 结构体实例
		account := &twitterAccount{
			Username: c.String("Username"),
			Password: c.String("Password"),
			Code:     c.String("Code"),
		}

		// 调用 LoginAndSaveCookies 方法
		err := account.LoginAndSaveCookiesTwoFactor()
		if err != nil {
			return err
		}

		return nil
	},
}

var Test = &cli.Command{
	Name:    "test",
	Aliases: []string{"t"},
	Usage:   "test",
	Action: func(c *cli.Context) error {
		println("test")
		return nil
	},
}

func writeSliceToCSV(filename string, data []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write([]string{value})
		if err != nil {
			return err
		}
	}

	return nil
}
