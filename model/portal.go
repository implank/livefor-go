package model

import "time"

type Greenbird struct {
	Order   int    `gorm:"primary_key" json:"order"`
	Title   string `gorm:"type:longtext" json:"title"`
	Content string `gorm:"type:longtext" json:"content"`
}

//createpost 0
//createcomment 1
//read post 2
//levelup 3
//ReadGreenbird 4
//Complete user info 5
//Upload user avatar 6
type SysMessage struct {
	UserID     uint64    `json:"user_id"`
	CreateTime time.Time `gorm:"default:Now()" json:"create_time"`
	Date       string    `json:"date"`
	Type       int       `json:"type"`
	Times      int       `gorm:"default:0" json:"got_exp"`
}

//LikePost 0
//LikeCommnet 1
//CreateComment 2
type Notification struct {
	UserID     uint64    `json:"user_id"`
	CreateTime time.Time `gorm:"default:Now()" json:"create_time"`
	Username   string    `grom:"size:255" json:"username"`
	PostID     uint64    `gorm:"not null" json:"post_id"`
	Title      string    `gorm:"type:longtext" json:"content"`
	Type       int       `json:"type"`
}

//api
type GreenbirdData struct {
	Greenbirds []Greenbird `json:"greenBirds"`
}
