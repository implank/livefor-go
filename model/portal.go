package model

import "time"

type Greenbird struct {
	Order   int    `gorm:"primary_key" json:"order"`
	Title   string `gorm:"type:text" json:"title"`
	Content string `gorm:"type:text" json:"content"`
}
type Notification struct {
	UserID   uint64 `gorm:"primary_key" json:"user_id"`
	CreateAt time.Time
}

//api
type GreenbirdData struct {
	Greenbirds []Greenbird `json:"greenBirds"`
}
