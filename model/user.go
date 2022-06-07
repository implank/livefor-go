package model

import "time"

type User struct {
	UserID    uint64    `gorm:"primary_key" json:"user_id"`
	Username  string    `gorm:"size:255;not null" json:"user_name"`
	Password  string    `gorm:"size:255;not null" json:"password"`
	Email     string    `gorm:"size:255;not null" json:"email"`
	UserLevel uint64    `gorm:"not null;default:1" json:"user_level"`
	Exp       uint64    `gorm:"not null;default:0" json:"exp"`
	Ban       bool      `gorm:"default:false" json:"ban"`
	Bandate   time.Time `gorm:"default:Now()" json:"bandate"`
	Sex       string    `gorm:"size:255" json:"sex"`
	Age       uint64    `gorm:"default:0" json:"age"`
	AvatarUrl string    `gorm:"default:null" json:"avatar_url"`
}
