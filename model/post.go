package model

import "time"

type Post struct {
	PostID          uint64    `gorm:"primary_key" json:"post_id"`
	UserID          uint64    `gorm:"not null" json:"user_id"`
	Username        string    `gorm:"size:255 not null" json:"username"`
	Title           string    `gorm:"size:255" json:"title"`
	Content         string    `gorm:"type:longtext" json:"content"`
	Views           uint64    `gorm:"default:0" json:"views"`
	Like            uint64    `gorm:"default:0" json:"like"`
	Dislike         uint64    `gorm:"default:0" json:"dislike"`
	MaxLike         uint64    `gorm:"default:0"`
	Comment         uint64    `gorm:"default:0" json:"comment"`
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
	Content     string    `gorm:"type:longtext" json:"content"`
	OnTop       bool      `gorm:"default:false" json:"on_top"`
	Like        uint64    `gorm:"default:0" json:"like"`
	Dislike     uint64    `gorm:"default:0" json:"dislike"`
	MaxLike     uint64    `gorm:"default:0"`
}
type CommentLike struct {
	CommentID     uint64 `gorm:"primary_key;auto_increment:false;" json:"comment_id"`
	UserID        uint64 `gorm:"primary_key;auto_increment:false;" json:"user_id"`
	LikeOrDislike bool   `gorm:"not null" json:"like_or_dislike"`
}
type SectionTag struct {
	Name    string `gorm:"primary_key;size:255;auto_increment:false" json:"name"`
	Section uint64 `gorm:"primary_key;default:0;auto_increment:false" json:"section"`
	Infers  uint64 `gorm:"default:0" json:"infers"`
}
type PostTag struct {
	PostID uint64 `gorm:"primary_key;auto_increment:false" json:"post_id"`
	Name   string `gorm:"primary_key;size:255" json:"name"`
}

//api struct
type Tag struct {
	Name string `json:"name"`
}
type CreatePostData struct {
	UserID  uint64 `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Section uint64 `json:"section"`
	Tags    []Tag  `json:"tags"`
}
type LikeCommentData CommentLike
type LikePostData PostLike
type GetPostsData struct {
	UserID  uint64 `json:"user_id"`
	Offset  uint64 `json:"offset"`
	Length  uint64 `json:"length"`
	Section uint64 `json:"section"`
	Order   string `json:"order"`
	Tags    []Tag  `json:"tags"`
}
type SearchPostsData struct {
	UserID  uint64   `json:"user_id"`
	Offset  uint64   `json:"offset"`
	Length  uint64   `json:"length"`
	Section uint64   `json:"section"`
	Order   string   `json:"order"`
	Fliters []string `json:"fliters"`
}
