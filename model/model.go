package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string `gorm:"uniqueIndex;not null" json:"username"`
	DisplayName   string `gorm:"not null" json:"display_name"`
	Email         string `gorm:"uniqueIndex;not null" json:"email"`
	PhoneNumber   string `json:"phone_number"`
	Password      string `gorm:"not null" json:"-"`
	EmailVerified bool   `gorm:"default:false" json:"email_verified"`
	PhoneVerified bool   `gorm:"default:true" json:"phone_verified"`
}
