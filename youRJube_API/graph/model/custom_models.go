package model

import (
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type Video struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	User        *User     `json:"user" gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	LocationID  string    `json:"location_id"`
	Location    *Location `json:"location" gorm:"ForeignKey:LocationID;AssociationForeignKey:ID"`
	Title       string    `json:"title"`
	View        int       `json:"view"`
	Thumbnail   string    `json:"thumbnail"`
	Link        string    `json:"link"`
	Description *string   `json:"description"`
	Restriction string    `json:"restriction"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type User struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	ProfilePict       string    `json:"profile_pict"`
	ChannelBackground string    `json:"channel_background"`
	Premium           bool      `json:"premium"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
type Location struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Comment struct {
	ID        string    `json:"id"`
	VideoID   string    `json:"video_id"`
	UserID    string    `json:"user_id"`
	ParentID  string    `json:"parent_id"`
	User      *User     `json:"user" gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	Comment   string    `json:"comment"`
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
