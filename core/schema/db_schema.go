package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

//go:generate go run entgo.io/ent/cmd/ent generate ./schema

// 整理好的 可以发布到nostr的内容. (可以配置发布计划,发布间隔等. (可以避免识别为bot账号, 被relay限速等. 或者nostr支持bot关联的协议之后,标记为bot账号))

type (
	TwitterProfile struct{ ent.Schema }
	TwitterNote    struct{ ent.Schema }
)

// TwitterFetchHistory 存储profile获取历史等,
// 可以用来追踪随时间变化的关注者数量, 历史签名, 历史头像等信息, 也算是可能重复获取的profile没有浪费
type TwitterFetchHistory struct{ ent.Schema }

type TwitterNotesFetchHistory struct{ ent.Schema }

type NostrNotesFetchHistory struct{ ent.Schema }

type BotEvents struct{ ent.Schema }

type Events struct{ ent.Schema }

func (TwitterProfile) Fields() []ent.Field {
	return []ent.Field{
		// required
		field.String("user_id").Unique().Comment("数字型唯一id, 如: 1687343259759349761"),
		field.String("user_name").Unique().Comment("用户可见的 唯一id: 如: Chai20230817 可修改,同一时间点不能重复"),
		field.String("name").Comment("显示名 如 柴静"),
		field.String("raw_yaml").Comment("raw profile data"),
		field.Bool("active").Optional().Comment("内部字段,是否继续更新此profile所对应的推文"),

		field.Time("joined_time").Optional().Comment("joined field on profile"),
		field.Int("likes_count").Optional().Comment("like count"),
		field.Int("listed_count").Optional().Comment("listed count"),
		field.Int("followers_count").Optional().Comment(""),
		field.Int("following_count").Optional().Comment(""),
		field.Int("friends_count").Optional().Comment(""),
		field.Int("twitter_count").Optional().Comment(""),
		field.String("pinned_note_id").Optional().Comment("原字段是多个值,这里只取第一个吧"),

		field.String("avatar").Optional().Comment(""),
		field.String("banner").Optional().Comment(""),
		field.String("biography").Optional().Comment(""),
		field.String("birthday").Optional().Comment(""),
		field.String("location").Optional().Comment(""),
		field.String("url").Optional().Comment(""),
		field.String("website").Optional().Comment(""),

		field.String("nsec_key").Unique().Optional(),
		field.String("npub_key").Optional(),
		field.String("private_key_hex").Optional(),
		field.String("public_key_hex").Optional(),
		field.Bool("publish_nostr").Default(false).Optional(),
	}
}

func (TwitterNote) Fields() []ent.Field {
	return []ent.Field{
		// required
		field.String("id").Unique().Comment("note unique id"),

		field.String("user_id").Comment("author id, (先不使用外键约束)"),
		field.String("user_name").Comment("user name"),
		field.String("conversation_id").Comment("conversation id"),
		field.String("text").Comment("原文形式的推文"),

		field.Int("likes").Comment("like count"),
		field.Int("replies").Comment("replies count"),
		field.Int("retweets").Comment("retweets count"),
		field.Int("views").Comment("view count"),

		field.Strings("photo_ids").Optional().Comment("photo ids"),
		field.Strings("photo_urls").Optional().Comment("photo urls"),

		field.Strings("video_id").Optional().Comment("video urls"),
		field.Strings("video_preview").Optional().Comment("video preview"),
		field.Strings("video_urls").Optional().Comment("video urls"),

		field.Strings("mention_id").Optional().Comment("mention id"),
		field.Strings("mention_username").Optional().Comment("mention username"),
		field.Strings("mention_name").Optional().Comment("mention name"),

		field.String("retweeted_status_id").Optional().Comment("retweeted target"),
		field.Bool("is_retweet").Optional(),

		field.Time("post_at").Comment("post time").Optional(),
		field.Strings("urls").Optional(),

		field.String("html_text").Optional().Comment("content in html format"),
		field.String("permanent_url").Optional(),
		field.Bool("is_reply").Optional(),
		field.Bool("is_pin").Optional(),
		field.Bool("sensitive_content").Optional(),
		field.String("reply_to").Optional(),
		field.String("raw_yaml").Comment("raw twitter data in yaml format"),
		field.Strings("hash_tags").Comment("twitter tags"),

		field.String("event").Optional().Comment("nostr event"),
		field.String("event_id").Optional().Comment("event id"),
		field.String("publish_status").Optional().Comment("发布结果: relay: status map"),
		field.Time("published_at").Optional().Comment("成功发布时间"),

		field.Time("fetched_at").Optional().Comment("note 初次获取时间,重复获取时此字段不更新"),
		field.String("language").Default("en").Comment("note language"),
	}
}

func (TwitterFetchHistory) Fields() []ent.Field {
	return []ent.Field{
		field.String("item_id").Comment("profile id, twitter id, ..."),
		field.String("name").Comment("readable name"),
		field.String("fetch_type").Comment("profile twitter ..."),
		field.Time("fetch_at").Default(time.Now),
		field.String("content").Comment("fetch result in yaml format"),
	}
}

func (TwitterNotesFetchHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment("auto inc id"),
		field.String("user_id").Comment("twitter user name"),
		field.String("username").Optional().Comment("twitter user name. 冗余字段,只是为了方便没有工具时查看, 减少join 操作."),
		field.String("name").Optional().Comment("twitter user name. 冗余字段,只是为了方便没有工具时查看, 减少join 操作."),
		field.Time("fetched_at").Comment("fetched time"),
		field.Int("fetch_count").Comment("fetched count"),
		field.Int("new_count").Comment("new count"),
	}
}

func (NostrNotesFetchHistory) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(100).
			Unique().
			Immutable(),
		field.Time("publish_at").Comment("post time"),
		field.String("public_key_hex").Comment("public key"),
		field.String("context").Comment("nostr context"),
	}
}

func (BotEvents) Fields() []ent.Field {
	return []ent.Field{
		field.String("event_id"),
		field.Time("inserted_at").Optional(),
	}
}

func (Events) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Bytes("event_id"),
		field.Bytes("pubkey"),
		field.Time("created_at"),
		field.Time("deleted_at").Optional().Nillable(),
		field.Int64("kind"),
		field.JSON("tags", [][]string{}).Default([][]string{}),
		field.String("content").Default(""),
		field.Bytes("sig"),
	}
}
