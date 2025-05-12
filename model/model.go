package model

import "time"

type User struct {
	ID            int64     `gorm:"primaryKey"`
	Username      string    `gorm:"uniqueIndex;not null" json:"username"`
	DisplayName   string    `gorm:"not null" json:"display_name"`
	Email         string    `gorm:"uniqueIndex;not null" json:"email"`
	PhoneNumber   string    `json:"phone_number"`
	Password      string    `gorm:"not null" json:"-"`
	EmailVerified bool      `gorm:"default:false" json:"email_verified"`
	PhoneVerified bool      `gorm:"default:true" json:"phone_verified"`
	CreatedAt     time.Time `json:"created_at"`
}
