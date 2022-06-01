package model

import "time"

type Post struct {
	PostID          uint64    `gorm:"primary_key" json:"post_id"`
	UserID          uint64    `gorm:"not null" json:"user_id"`
	Title           string    `gorm:"size:255" json:"title"`
	Content         string    `gorm:"type:text" json:"content"`
	Views           uint64    `gorm:"default:0" json:"views"`
	Like            uint64    `gorm:"default:0" json:"like"`
	Dislike         uint64    `gorm:"default:0" json:"dislike"`
	Section         uint64    `gorm:"default:0" json:"section"`
	CreateTime      time.Time `gorm:"default:Now()" json:"create_time"`
	LastCommentTime time.Time `gorm:"default:Now()" json:"last_comment_time"`
}
type PostLike struct {
	PostID        uint64 `gorm:"primary_key;auto_increment:false" json:"post_id"`
	UserID        uint64 `gorm:"primary_key;auto_increment:false" json:"user_id"`
	LikeOrDislike bool   `gorm:"not null" json:"like_or_dislike"`
}
type Comment struct {
	CommentID   uint64    `gorm:"primary_key;" json:"comment_id"`
	UserID      uint64    `gorm:"not null" json:"user_id"`
	Username    string    `gorm:"size:255 not null" json:"username"`
	PostID      uint64    `gorm:"not null" json:"post_id"`
	CommentTime time.Time `json:"comment_time"`
	Content     string    `gorm:"size:255 not null" json:"content"`
	OnTop       bool      `gorm:"default:false" json:"on_top"`
	Like        uint64    `gorm:"default:0" json:"like"`
	Dislike     uint64    `gorm:"default:0" json:"dislike"`
}
type CommentLike struct {
	CommentID     uint64 `gorm:"primary_key;auto_increment:false;" json:"comment_id"`
	UserID        uint64 `gorm:"primary_key;auto_increment:false;" json:"user_id"`
	LikeOrDislike bool   `gorm:"not null" json:"like_or_dislike"`
}
type Tag struct {
	Name string `gorm:"primary_key;size:255" json:"name"`
}
type PostTag struct {
	PostID uint64 `gorm:"primary_key;auto_increment:false" json:"post_id"`
	Name   string `gorm:"primary_key;size:255" json:"name"`
}

//api struct
type CreatePostData struct {
	UserID  uint64 `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Section uint64 `json:"section"`
	Tags    []Tag  `json:"tags"`
}
type LikeCommentData CommentLike
