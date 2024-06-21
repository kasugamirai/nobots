package nostr_fixture

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"freefrom.space/nobot/biz/dal/postgres"
	"freefrom.space/nobot/biz/dal/sqlite"
	"freefrom.space/nobot/core"
	"freefrom.space/nobot/core/nostrnotesfetchhistory"
	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
	"github.com/rs/zerolog/log"
)

// PublishSignedEventToRelay sends a signed event to multiple relay servers for publishing
//
// Parameters:
//   - ctx: the context.Context for the function
//   - relayPool: the pool of relay servers
//   - relayURLList: the list of relay server URLs to publish to
//   - event: the event to publish
//
// Returns:
//   - bool: whether at least one relay server successfully published the event
//   - error: any error that occurred during publishing
func PublishSignedEventToRelay(ctx context.Context,
	relayPool *nostr.SimplePool,
	relayURLList []string, event *nostr.Event,
) (bool, error) {
	total := len(relayURLList)
	if total == 0 {
		return false, errors.New("target relay list empty")
	}
	okCount := 0
	failedCount := 0
	for _, relayURL := range relayURLList {
		relay, err := relayPool.EnsureRelay(relayURL)
		if err != nil {
			log.Err(err).
				Str("relay", relayURL).
				Msg("relay ensure connection failed")
			failedCount += 1
			continue
		}
		if err := relay.Publish(ctx, *event); err != nil {
			log.Err(err).
				Str("relay", relayURL).
				Msg("relay publish failed")
			failedCount += 1
		} else {
			okCount += 1
		}
	}
	return okCount > 0, nil
}

// QueryDataFromRelay retrieves data from multiple relay servers based on the provided filters.
//
// Parameters:
//   - ctx: the context.Context for the function
//   - relayPool: the pool of relay servers
//   - relayURLList: the list of relay server URLs to query from
//   - filters: the filters to apply while querying data
//
// Returns:
//   - []*nostr.Event: the list of events that match the provided filters
//   - error: any error that occurred during querying
func QueryDataFromRelay(ctx context.Context,
	relayPool *nostr.SimplePool,
	relayURLList []string, filters nostr.Filter,
) ([]*nostr.Event, error) {
	total := len(relayURLList)
	if total == 0 {
		return nil, errors.New("target relay list empty")
	}
	var events []*nostr.Event
	for _, relayURL := range relayURLList {
		relay, err := relayPool.EnsureRelay(relayURL)
		if err != nil {
			log.Err(err).
				Str("relay", relayURL).
				Msg("relay ensure connection failed")
			continue
		}
		queriedEvents, err := relay.QuerySync(ctx, filters)
		if err != nil {
			log.Err(err).
				Str("relay", relayURL).
				Msg("relay query failed")
			continue
		}
		events = append(events, queriedEvents...)
	}
	return events, nil
}

func CloseRelayPool(relayPool *nostr.SimplePool) {
	relayPool.Relays.Range(func(url string, relay *nostr.Relay) bool {
		fmt.Printf("Closing relay at URL: %s\n", url)
		if relay.IsConnected() {
			err := relay.Close()
			if err != nil {
				log.Err(err).
					Str("relay", url).
					Msg("failed to close relay")
			}
		}
		return true
	})
}

func ConvertTwitterContentToNostrEvent(note *core.TwitterNote, profile *core.TwitterProfile) *nostr.Event {
	content := note.Text

	for _, photoURL := range note.PhotoUrls {
		content += " " + photoURL
	}
	for _, videoURL := range note.VideoUrls {
		content += " " + videoURL
	}

	anime := nostr.Tag{"t", "anime"}
	acg := nostr.Tag{"t", "acg"}
	git := nostr.Tag{"t", "https://github.com/kasugamirai/nobots"}
	waifu := nostr.Tag{"t", "waifu"}
	tags := []nostr.Tag{anime, acg, git, waifu}

	//rand.Seed(time.Now().UnixNano())
	//randomMinutes := time.Duration(rand.Intn(31)) * time.Minute

	//adjustedTime := note.PostAt.Add(-randomMinutes)

	//timestamp := nostr.Timestamp(adjustedTime.Unix())

	textEvent := &nostr.Event{
		CreatedAt: nostr.Timestamp(time.Now().Unix()),
		//CreatedAt: nostr.Timestamp(timestamp),
		Kind:    nostr.KindTextNote,
		Tags:    tags,
		Content: content,
	}

	_ = textEvent.Sign(profile.PrivateKeyHex)

	log.Info().
		Str("e", textEvent.String()).
		Msg("event")

	return textEvent
}

func CatchUserDataFromRelay(relayList []string, pubKeyHex []string) error {
	client := sqlite.Client
	postgreClient := postgres.Client
	defer postgreClient.Close()
	defer client.Close()
	for _, pub := range pubKeyHex {
		log.Info().
			Str("pub", pub).
			Msg("start")
		ctx := context.Background()
		var filter nostr.Filter
		if _, v, err := nip19.Decode(pub); err == nil {
			pub := v.(string)
			filter = nostr.Filter{
				Kinds:   []int{nostr.KindTextNote},
				Authors: []string{pub},
			}
		} else {
			log.Err(err).
				Msg("failed to decode pub key")
		}

		events, err := QueryDataFromRelay(ctx, nostr.NewSimplePool(ctx), relayList, filter)

		CloseRelayPool(nostr.NewSimplePool(ctx))

		if err != nil {
			log.Err(err).Msg("failed to query data from relay")
		}

		// Iterate over the received events.
		for _, ev := range events {
			select {
			case <-ctx.Done():
				break
			default:
				exists, err := client.NostrNotesFetchHistory.
					Query().
					Where(nostrnotesfetchhistory.IDEQ(ev.ID)).
					Exist(ctx)
				if exists {
					log.Info().Msg("event exists")
					continue
				}
				
				// Convert Unix time to time.Time object.
				unixTimestamp := int64(ev.CreatedAt)
				eventTime := time.Unix(unixTimestamp, 0)

				// Create and save a new Event record in the database.
				event, err := client.NostrNotesFetchHistory.
					Create().
					SetPublicKeyHex(ev.PubKey).
					SetContext(ev.Content).
					SetPublishAt(eventTime).
					SetID(ev.ID).
					Save(ctx)
				if err != nil {
					log.Err(err).Msg("failed to create event")
					continue
				}
				log.Info().Interface("event", event).Msg("event created")
			}
			log.Info().Msg("Events stored in the database successfully.")
		}
	}

	return nil
}

// filter replay message 沈老师不要回复的信息
func containsReply(tags []nostr.Tag) bool {
	for _, tag := range tags {
		for _, t := range tag {
			if strings.Contains(t, "reply") {
				return true
			}
		}
	}
	return false
}

func GetUserFollowings(relayList []string, pubKeyHex string) ([]string, error) {
	ctx := context.Background()
	var filter nostr.Filter
	if _, v, err := nip19.Decode(pubKeyHex); err == nil {
		pub := v.(string)
		filter = nostr.Filter{
			Kinds:   []int{nostr.KindContactList},
			Authors: []string{pub},
		}
	} else {
		log.Err(err).
			Msg("failed to decode pub key")
		return nil, err
	}

	events, err := QueryDataFromRelay(ctx, nostr.NewSimplePool(ctx), relayList, filter)
	if err != nil {
		log.Err(err).Msg("failed to query data from relay")
		return nil, err
	}

	// Sort events by CreatedAt in descending order
	sort.SliceStable(events, func(i, j int) bool {
		return events[i].CreatedAt > events[j].CreatedAt
	})

	// Only keep the latest event
	latestEvent := events[0]
	var followings []string

	for _, tag := range latestEvent.Tags {
		if tag.Key() == "p" {
			npub, err := nip19.EncodePublicKey(tag.Value())
			if err != nil {
				log.Err(err).Msg("failed to encode public key")
				continue
			}
			followings = append(followings, npub)
		}
	}

	return followings, nil
}
