package model

import "time"

type Greenbird struct {
	Order   int    `gorm:"primary_key" json:"order"`
	Title   string `gorm:"type:longtext" json:"title"`
	Content string `gorm:"type:longtext" json:"content"`
}
type Notification struct {
	UserID     uint64    `json:"user_id"`
	CreateTime time.Time `gorm:"default:Now()" json:"create_time"`
	Username   string    `grom:"size:255" json:"username"`
	PostID     uint64    `gorm:"not null" json:"post_id"`
	Msg        string    `gorm:"type:text" json:"msg"`
	Content    string    `gorm:"type:longtext" json:"content"`
}

//api
type GreenbirdData struct {
	Greenbirds []Greenbird `json:"greenBirds"`
}
