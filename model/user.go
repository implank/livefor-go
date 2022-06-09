package model

import "time"

type User struct {
	UserID   uint64 `gorm:"primary_key" json:"user_id"`
	Username string `gorm:"size:255;not null" json:"user_name"`
	//UserInfo  string    `gorm:"type:text" json:"user_info"`
	Password      string    `gorm:"size:255;not null" json:"password"`
	Email         string    `gorm:"size:255;not null" json:"email"`
	Level         int       `gorm:"not null;default:1" json:"user_level"`
	Exp           int       `gorm:"not null;default:0" json:"exp"`
	Ban           bool      `gorm:"default:false" json:"ban"`
	Bandate       time.Time `gorm:"default:Now()" json:"bandate"`
	Sex           string    `gorm:"size:255;default:'未知'" json:"sex"`
	Age           uint64    `gorm:"default:0" json:"age"`
	AvatarUrl     string    `gorm:"default:http://43.138.77.133:81/media/avatars/default.jpg" json:"avatar_url"`
	Confirmed     bool      `gorm:"default:false" json:"confirmed"`
	ConfirmNumber int       `gorm:"default:0" json:"confirmed_number"`
}
