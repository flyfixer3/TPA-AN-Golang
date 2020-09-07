package model

import (
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type User struct {
	ID                string        `json:"id" gorm:"primaryKey"`
	Name              string        `json:"name"`
	Email             string        `json:"email"`
	ProfilePict       string        `json:"profile_pict"`
	ChannelBackground string        `json:"channel_background"`
	Premium           bool          `json:"premium"`
	Subscribers       []*Subscriber `json:"subscribers" gorm:"foreignKey:SubscriberID;AssociationForeignKey:ID"`
	SavePlaylists     []*Playlist   `json:"save_playlists" gorm:"many2many:save_playlists;"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
}

type Subscriber struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	TargetID     string    `json:"target_id" gorm:"primaryKey"`
	Target       *User     `json:"target" gorm:"ForeignKey:TargetID;AssociationForeignKey:ID"`
	SubscriberID string    `json:"subscriber_id" gorm:"primaryKey"`
	Notification bool      `json:"notification"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Playlist struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	UserID      string    `json:"user_id"`
	User        *User     `json:"user" gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Private     bool      `json:"private"`
	Videos      []*Video  `json:"videos" gorm:"many2many:playlist_details"`
	View        int       `json:"view"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type Video struct {
	ID          string       `json:"id"`
	UserID      string       `json:"user_id"`
	User        *User        `json:"user" gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	LocationID  string       `json:"location_id"`
	Location    *Location    `json:"location" gorm:"ForeignKey:LocationID;AssociationForeignKey:ID"`
	LikeVideo   []*LikeVideo `json:"like_video" gorm:"ForeignKey:VideoID;AssociationForeignKey:ID"`
	Title       string       `json:"title"`
	View        int          `json:"view"`
	Thumbnail   string       `json:"thumbnail"`
	Link        string       `json:"link"`
	Description *string      `json:"description"`
	Restriction string       `json:"restriction"`
	Duration    int          `json:"duration"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

type PlaylistDetail struct {
	PlaylistID string    `json:"playlist_id"`
	VideoID    string    `json:"video_id"`
	Priority   int       `json:"priority"`
	Video      *Video    `json:"video" gorm:"ForeignKey:VideoID;AssociationForeignKey:ID"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
type SavePlaylist struct {
	PlaylistID string `json:"playlist_id"`
	UserID     string `json:"User_id"`
}

type Location struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Comment struct {
	ID          string         `json:"id"`
	VideoID     string         `json:"video_id"`
	UserID      string         `json:"user_id"`
	ParentID    string         `json:"parent_id"`
	User        *User          `json:"user" gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	ReplyTo     string         `json:"reply_to"`
	Comment     string         `json:"comment"`
	LikeComment []*LikeComment `json:"like_comment" gorm:"ForeignKey:CommentID;AssociationForeignKey:ID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type Post struct {
	ID          string      `json:"id"`
	UserID      string      `json:"user_id"`
	User        *User       `json:"user" gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	Title       string      `json:"title"`
	Description *string     `json:"description"`
	Media       *string     `json:"status"`
	LikePost    []*LikePost `json:"like_post" gorm:"ForeignKey:PostID;AssociationForeignKey:ID"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type LikeComment struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	CommentID string    `json:"comment_id"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LikeVideo struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	VideoID   string    `json:"video_id"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type LikePost struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	PostID    string    `json:"video_id"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var (
	ServerError         = GenerateError("Something went wrong! Please try again later")
	UserNotExist        = GenerateError("User not exists")
	UnauthorisedError   = GenerateError("You are not authorised to perform this action")
	TimeStampError      = GenerateError("time should be a unix timestamp")
	InternalServerError = GenerateError("internal server error")
)

func MarshalTimestamp(t time.Time) graphql.Marshaler {
	timestamp := t.Unix() * 1000

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(timestamp, 10))
	})
}

func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(int); ok {
		return time.Unix(int64(tmpStr), 0), nil
	}
	return time.Time{}, TimeStampError
}
func GenerateError(err string) error {
	return errors.New(err)
}
