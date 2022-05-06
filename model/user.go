package model

type User struct {
	UserID    uint64 `gorm:"primary_key" json:"user_id"`
	Username  string `gorm:"size:255 not null" json:"user_name"`
	Password  string `gorm:"size:255 not null" json:"password"`
	Email     string `gorm:"size:255 not null" json:"email"`
	UserLevel uint64 `gorm:"not null default:0" json:"user_level"`
	Exp       uint64 `gorm:"not null default:0" json:"exp"`
}
