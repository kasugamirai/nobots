// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BotEventsColumns holds the columns for the "bot_events" table.
	BotEventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "event_id", Type: field.TypeString},
		{Name: "inserted_at", Type: field.TypeTime, Nullable: true},
	}
	// BotEventsTable holds the schema information for the "bot_events" table.
	BotEventsTable = &schema.Table{
		Name:       "bot_events",
		Columns:    BotEventsColumns,
		PrimaryKey: []*schema.Column{BotEventsColumns[0]},
	}
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "event_id", Type: field.TypeBytes},
		{Name: "pubkey", Type: field.TypeBytes},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "kind", Type: field.TypeInt64},
		{Name: "tags", Type: field.TypeJSON},
		{Name: "content", Type: field.TypeString, Default: ""},
		{Name: "sig", Type: field.TypeBytes},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
	}
	// NostrNotesFetchHistoriesColumns holds the columns for the "nostr_notes_fetch_histories" table.
	NostrNotesFetchHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 100},
		{Name: "publish_at", Type: field.TypeTime},
		{Name: "public_key_hex", Type: field.TypeString},
		{Name: "context", Type: field.TypeString},
	}
	// NostrNotesFetchHistoriesTable holds the schema information for the "nostr_notes_fetch_histories" table.
	NostrNotesFetchHistoriesTable = &schema.Table{
		Name:       "nostr_notes_fetch_histories",
		Columns:    NostrNotesFetchHistoriesColumns,
		PrimaryKey: []*schema.Column{NostrNotesFetchHistoriesColumns[0]},
	}
	// TwitterFetchHistoriesColumns holds the columns for the "twitter_fetch_histories" table.
	TwitterFetchHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "item_id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "fetch_type", Type: field.TypeString},
		{Name: "fetch_at", Type: field.TypeTime},
		{Name: "content", Type: field.TypeString},
	}
	// TwitterFetchHistoriesTable holds the schema information for the "twitter_fetch_histories" table.
	TwitterFetchHistoriesTable = &schema.Table{
		Name:       "twitter_fetch_histories",
		Columns:    TwitterFetchHistoriesColumns,
		PrimaryKey: []*schema.Column{TwitterFetchHistoriesColumns[0]},
	}
	// TwitterNotesColumns holds the columns for the "twitter_notes" table.
	TwitterNotesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "user_id", Type: field.TypeString},
		{Name: "user_name", Type: field.TypeString},
		{Name: "conversation_id", Type: field.TypeString},
		{Name: "text", Type: field.TypeString},
		{Name: "likes", Type: field.TypeInt},
		{Name: "replies", Type: field.TypeInt},
		{Name: "retweets", Type: field.TypeInt},
		{Name: "views", Type: field.TypeInt},
		{Name: "photo_ids", Type: field.TypeJSON, Nullable: true},
		{Name: "photo_urls", Type: field.TypeJSON, Nullable: true},
		{Name: "video_id", Type: field.TypeJSON, Nullable: true},
		{Name: "video_preview", Type: field.TypeJSON, Nullable: true},
		{Name: "video_urls", Type: field.TypeJSON, Nullable: true},
		{Name: "mention_id", Type: field.TypeJSON, Nullable: true},
		{Name: "mention_username", Type: field.TypeJSON, Nullable: true},
		{Name: "mention_name", Type: field.TypeJSON, Nullable: true},
		{Name: "retweeted_status_id", Type: field.TypeString, Nullable: true},
		{Name: "is_retweet", Type: field.TypeBool, Nullable: true},
		{Name: "post_at", Type: field.TypeTime, Nullable: true},
		{Name: "urls", Type: field.TypeJSON, Nullable: true},
		{Name: "html_text", Type: field.TypeString, Nullable: true},
		{Name: "permanent_url", Type: field.TypeString, Nullable: true},
		{Name: "is_reply", Type: field.TypeBool, Nullable: true},
		{Name: "is_pin", Type: field.TypeBool, Nullable: true},
		{Name: "sensitive_content", Type: field.TypeBool, Nullable: true},
		{Name: "reply_to", Type: field.TypeString, Nullable: true},
		{Name: "raw_yaml", Type: field.TypeString},
		{Name: "hash_tags", Type: field.TypeJSON},
		{Name: "event", Type: field.TypeString, Nullable: true},
		{Name: "event_id", Type: field.TypeString, Nullable: true},
		{Name: "publish_status", Type: field.TypeString, Nullable: true},
		{Name: "published_at", Type: field.TypeTime, Nullable: true},
		{Name: "fetched_at", Type: field.TypeTime, Nullable: true},
		{Name: "language", Type: field.TypeString, Default: "en"},
	}
	// TwitterNotesTable holds the schema information for the "twitter_notes" table.
	TwitterNotesTable = &schema.Table{
		Name:       "twitter_notes",
		Columns:    TwitterNotesColumns,
		PrimaryKey: []*schema.Column{TwitterNotesColumns[0]},
	}
	// TwitterNotesFetchHistoriesColumns holds the columns for the "twitter_notes_fetch_histories" table.
	TwitterNotesFetchHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "user_id", Type: field.TypeString},
		{Name: "username", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "fetched_at", Type: field.TypeTime},
		{Name: "fetch_count", Type: field.TypeInt},
		{Name: "new_count", Type: field.TypeInt},
	}
	// TwitterNotesFetchHistoriesTable holds the schema information for the "twitter_notes_fetch_histories" table.
	TwitterNotesFetchHistoriesTable = &schema.Table{
		Name:       "twitter_notes_fetch_histories",
		Columns:    TwitterNotesFetchHistoriesColumns,
		PrimaryKey: []*schema.Column{TwitterNotesFetchHistoriesColumns[0]},
	}
	// TwitterProfilesColumns holds the columns for the "twitter_profiles" table.
	TwitterProfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeString, Unique: true},
		{Name: "user_name", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "raw_yaml", Type: field.TypeString},
		{Name: "active", Type: field.TypeBool, Nullable: true},
		{Name: "joined_time", Type: field.TypeTime, Nullable: true},
		{Name: "likes_count", Type: field.TypeInt, Nullable: true},
		{Name: "listed_count", Type: field.TypeInt, Nullable: true},
		{Name: "followers_count", Type: field.TypeInt, Nullable: true},
		{Name: "following_count", Type: field.TypeInt, Nullable: true},
		{Name: "friends_count", Type: field.TypeInt, Nullable: true},
		{Name: "twitter_count", Type: field.TypeInt, Nullable: true},
		{Name: "pinned_note_id", Type: field.TypeString, Nullable: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true},
		{Name: "banner", Type: field.TypeString, Nullable: true},
		{Name: "biography", Type: field.TypeString, Nullable: true},
		{Name: "birthday", Type: field.TypeString, Nullable: true},
		{Name: "location", Type: field.TypeString, Nullable: true},
		{Name: "url", Type: field.TypeString, Nullable: true},
		{Name: "website", Type: field.TypeString, Nullable: true},
		{Name: "nsec_key", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "npub_key", Type: field.TypeString, Nullable: true},
		{Name: "private_key_hex", Type: field.TypeString, Nullable: true},
		{Name: "public_key_hex", Type: field.TypeString, Nullable: true},
		{Name: "publish_nostr", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// TwitterProfilesTable holds the schema information for the "twitter_profiles" table.
	TwitterProfilesTable = &schema.Table{
		Name:       "twitter_profiles",
		Columns:    TwitterProfilesColumns,
		PrimaryKey: []*schema.Column{TwitterProfilesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BotEventsTable,
		EventsTable,
		NostrNotesFetchHistoriesTable,
		TwitterFetchHistoriesTable,
		TwitterNotesTable,
		TwitterNotesFetchHistoriesTable,
		TwitterProfilesTable,
	}
)

func init() {
}
