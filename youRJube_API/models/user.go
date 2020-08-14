package models

import "time"

// User -> model user
type User struct {
	ID                string
	Name              string `json:"name" binding:"required"`
	Email             string `gorm:"unique" json:"email" binding:"required,email" `
	Password          string `json:"password" binding:"required,min=6"`
	ProfilePict       *string
	ChannelName       *string
	ChannelPict       *string
	ChannelBackground *string
	Premium           *bool `gorm:"default:FALSE"`
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}
