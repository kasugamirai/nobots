// Code generated by ent, DO NOT EDIT.

package core

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"freefrom.space/nobot/core/twitterprofile"
)

// TwitterProfile is the model entity for the TwitterProfile schema.
type TwitterProfile struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 数字型唯一id, 如: 1687343259759349761
	UserID string `json:"user_id,omitempty"`
	// 用户可见的 唯一id: 如: Chai20230817 可修改,同一时间点不能重复
	UserName string `json:"user_name,omitempty"`
	// 显示名 如 柴静
	Name string `json:"name,omitempty"`
	// raw profile data
	RawYaml string `json:"raw_yaml,omitempty"`
	// 内部字段,是否继续更新此profile所对应的推文
	Active bool `json:"active,omitempty"`
	// joined field on profile
	JoinedTime time.Time `json:"joined_time,omitempty"`
	// like count
	LikesCount int `json:"likes_count,omitempty"`
	// listed count
	ListedCount int `json:"listed_count,omitempty"`
	// FollowersCount holds the value of the "followers_count" field.
	FollowersCount int `json:"followers_count,omitempty"`
	// FollowingCount holds the value of the "following_count" field.
	FollowingCount int `json:"following_count,omitempty"`
	// FriendsCount holds the value of the "friends_count" field.
	FriendsCount int `json:"friends_count,omitempty"`
	// TwitterCount holds the value of the "twitter_count" field.
	TwitterCount int `json:"twitter_count,omitempty"`
	// 原字段是多个值,这里只取第一个吧
	PinnedNoteID string `json:"pinned_note_id,omitempty"`
	// Avatar holds the value of the "avatar" field.
	Avatar string `json:"avatar,omitempty"`
	// Banner holds the value of the "banner" field.
	Banner string `json:"banner,omitempty"`
	// Biography holds the value of the "biography" field.
	Biography string `json:"biography,omitempty"`
	// Birthday holds the value of the "birthday" field.
	Birthday string `json:"birthday,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Website holds the value of the "website" field.
	Website string `json:"website,omitempty"`
	// NsecKey holds the value of the "nsec_key" field.
	NsecKey string `json:"nsec_key,omitempty"`
	// NpubKey holds the value of the "npub_key" field.
	NpubKey string `json:"npub_key,omitempty"`
	// PrivateKeyHex holds the value of the "private_key_hex" field.
	PrivateKeyHex string `json:"private_key_hex,omitempty"`
	// PublicKeyHex holds the value of the "public_key_hex" field.
	PublicKeyHex string `json:"public_key_hex,omitempty"`
	// PublishNostr holds the value of the "publish_nostr" field.
	PublishNostr bool `json:"publish_nostr,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TwitterProfile) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case twitterprofile.FieldActive, twitterprofile.FieldPublishNostr:
			values[i] = new(sql.NullBool)
		case twitterprofile.FieldID, twitterprofile.FieldLikesCount, twitterprofile.FieldListedCount, twitterprofile.FieldFollowersCount, twitterprofile.FieldFollowingCount, twitterprofile.FieldFriendsCount, twitterprofile.FieldTwitterCount:
			values[i] = new(sql.NullInt64)
		case twitterprofile.FieldUserID, twitterprofile.FieldUserName, twitterprofile.FieldName, twitterprofile.FieldRawYaml, twitterprofile.FieldPinnedNoteID, twitterprofile.FieldAvatar, twitterprofile.FieldBanner, twitterprofile.FieldBiography, twitterprofile.FieldBirthday, twitterprofile.FieldLocation, twitterprofile.FieldURL, twitterprofile.FieldWebsite, twitterprofile.FieldNsecKey, twitterprofile.FieldNpubKey, twitterprofile.FieldPrivateKeyHex, twitterprofile.FieldPublicKeyHex:
			values[i] = new(sql.NullString)
		case twitterprofile.FieldJoinedTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TwitterProfile fields.
func (tp *TwitterProfile) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case twitterprofile.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tp.ID = int(value.Int64)
		case twitterprofile.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				tp.UserID = value.String
			}
		case twitterprofile.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_name", values[i])
			} else if value.Valid {
				tp.UserName = value.String
			}
		case twitterprofile.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				tp.Name = value.String
			}
		case twitterprofile.FieldRawYaml:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field raw_yaml", values[i])
			} else if value.Valid {
				tp.RawYaml = value.String
			}
		case twitterprofile.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				tp.Active = value.Bool
			}
		case twitterprofile.FieldJoinedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field joined_time", values[i])
			} else if value.Valid {
				tp.JoinedTime = value.Time
			}
		case twitterprofile.FieldLikesCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field likes_count", values[i])
			} else if value.Valid {
				tp.LikesCount = int(value.Int64)
			}
		case twitterprofile.FieldListedCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field listed_count", values[i])
			} else if value.Valid {
				tp.ListedCount = int(value.Int64)
			}
		case twitterprofile.FieldFollowersCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field followers_count", values[i])
			} else if value.Valid {
				tp.FollowersCount = int(value.Int64)
			}
		case twitterprofile.FieldFollowingCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field following_count", values[i])
			} else if value.Valid {
				tp.FollowingCount = int(value.Int64)
			}
		case twitterprofile.FieldFriendsCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field friends_count", values[i])
			} else if value.Valid {
				tp.FriendsCount = int(value.Int64)
			}
		case twitterprofile.FieldTwitterCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field twitter_count", values[i])
			} else if value.Valid {
				tp.TwitterCount = int(value.Int64)
			}
		case twitterprofile.FieldPinnedNoteID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pinned_note_id", values[i])
			} else if value.Valid {
				tp.PinnedNoteID = value.String
			}
		case twitterprofile.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				tp.Avatar = value.String
			}
		case twitterprofile.FieldBanner:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field banner", values[i])
			} else if value.Valid {
				tp.Banner = value.String
			}
		case twitterprofile.FieldBiography:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field biography", values[i])
			} else if value.Valid {
				tp.Biography = value.String
			}
		case twitterprofile.FieldBirthday:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field birthday", values[i])
			} else if value.Valid {
				tp.Birthday = value.String
			}
		case twitterprofile.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				tp.Location = value.String
			}
		case twitterprofile.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				tp.URL = value.String
			}
		case twitterprofile.FieldWebsite:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field website", values[i])
			} else if value.Valid {
				tp.Website = value.String
			}
		case twitterprofile.FieldNsecKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nsec_key", values[i])
			} else if value.Valid {
				tp.NsecKey = value.String
			}
		case twitterprofile.FieldNpubKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field npub_key", values[i])
			} else if value.Valid {
				tp.NpubKey = value.String
			}
		case twitterprofile.FieldPrivateKeyHex:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field private_key_hex", values[i])
			} else if value.Valid {
				tp.PrivateKeyHex = value.String
			}
		case twitterprofile.FieldPublicKeyHex:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field public_key_hex", values[i])
			} else if value.Valid {
				tp.PublicKeyHex = value.String
			}
		case twitterprofile.FieldPublishNostr:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field publish_nostr", values[i])
			} else if value.Valid {
				tp.PublishNostr = value.Bool
			}
		default:
			tp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TwitterProfile.
// This includes values selected through modifiers, order, etc.
func (tp *TwitterProfile) Value(name string) (ent.Value, error) {
	return tp.selectValues.Get(name)
}

// Update returns a builder for updating this TwitterProfile.
// Note that you need to call TwitterProfile.Unwrap() before calling this method if this TwitterProfile
// was returned from a transaction, and the transaction was committed or rolled back.
func (tp *TwitterProfile) Update() *TwitterProfileUpdateOne {
	return NewTwitterProfileClient(tp.config).UpdateOne(tp)
}

// Unwrap unwraps the TwitterProfile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tp *TwitterProfile) Unwrap() *TwitterProfile {
	_tx, ok := tp.config.driver.(*txDriver)
	if !ok {
		panic("core: TwitterProfile is not a transactional entity")
	}
	tp.config.driver = _tx.drv
	return tp
}

// String implements the fmt.Stringer.
func (tp *TwitterProfile) String() string {
	var builder strings.Builder
	builder.WriteString("TwitterProfile(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tp.ID))
	builder.WriteString("user_id=")
	builder.WriteString(tp.UserID)
	builder.WriteString(", ")
	builder.WriteString("user_name=")
	builder.WriteString(tp.UserName)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(tp.Name)
	builder.WriteString(", ")
	builder.WriteString("raw_yaml=")
	builder.WriteString(tp.RawYaml)
	builder.WriteString(", ")
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", tp.Active))
	builder.WriteString(", ")
	builder.WriteString("joined_time=")
	builder.WriteString(tp.JoinedTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("likes_count=")
	builder.WriteString(fmt.Sprintf("%v", tp.LikesCount))
	builder.WriteString(", ")
	builder.WriteString("listed_count=")
	builder.WriteString(fmt.Sprintf("%v", tp.ListedCount))
	builder.WriteString(", ")
	builder.WriteString("followers_count=")
	builder.WriteString(fmt.Sprintf("%v", tp.FollowersCount))
	builder.WriteString(", ")
	builder.WriteString("following_count=")
	builder.WriteString(fmt.Sprintf("%v", tp.FollowingCount))
	builder.WriteString(", ")
	builder.WriteString("friends_count=")
	builder.WriteString(fmt.Sprintf("%v", tp.FriendsCount))
	builder.WriteString(", ")
	builder.WriteString("twitter_count=")
	builder.WriteString(fmt.Sprintf("%v", tp.TwitterCount))
	builder.WriteString(", ")
	builder.WriteString("pinned_note_id=")
	builder.WriteString(tp.PinnedNoteID)
	builder.WriteString(", ")
	builder.WriteString("avatar=")
	builder.WriteString(tp.Avatar)
	builder.WriteString(", ")
	builder.WriteString("banner=")
	builder.WriteString(tp.Banner)
	builder.WriteString(", ")
	builder.WriteString("biography=")
	builder.WriteString(tp.Biography)
	builder.WriteString(", ")
	builder.WriteString("birthday=")
	builder.WriteString(tp.Birthday)
	builder.WriteString(", ")
	builder.WriteString("location=")
	builder.WriteString(tp.Location)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(tp.URL)
	builder.WriteString(", ")
	builder.WriteString("website=")
	builder.WriteString(tp.Website)
	builder.WriteString(", ")
	builder.WriteString("nsec_key=")
	builder.WriteString(tp.NsecKey)
	builder.WriteString(", ")
	builder.WriteString("npub_key=")
	builder.WriteString(tp.NpubKey)
	builder.WriteString(", ")
	builder.WriteString("private_key_hex=")
	builder.WriteString(tp.PrivateKeyHex)
	builder.WriteString(", ")
	builder.WriteString("public_key_hex=")
	builder.WriteString(tp.PublicKeyHex)
	builder.WriteString(", ")
	builder.WriteString("publish_nostr=")
	builder.WriteString(fmt.Sprintf("%v", tp.PublishNostr))
	builder.WriteByte(')')
	return builder.String()
}

// TwitterProfiles is a parsable slice of TwitterProfile.
type TwitterProfiles []*TwitterProfile
