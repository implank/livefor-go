package service

import (
	"gin-project/global"
	"gin-project/model"

	"github.com/jinzhu/gorm"
)

func CreatePost(post *model.Post) (err error) {
	if err = global.DB.Create(post).Error; err != nil {
		return err
	}
	return
}
func QueryPost(postID int64) (post model.Post, notFound bool) {
	notFound = global.DB.First(&post, postID).RecordNotFound()
	return post, notFound
}
func DeletePost(postID int64) (err error) {
	var post model.Post
	notFound := global.DB.First(&post, postID).RecordNotFound()
	if notFound {
		return gorm.ErrRecordNotFound
	}
	err = global.DB.Delete(&post).Error
	return err
}
func CreateComment(comment *model.Comment) (err error) {
	if err = global.DB.Create(comment).Error; err != nil {
		return err
	}
	return
}
func QueryComment(commentID uint64) (comment model.Comment, notFound bool) {
	notFound = global.DB.First(&comment, commentID).RecordNotFound()
	return comment, notFound
}
func DeleteComment(commentID uint64) (err error) {
	var comment model.Comment
	notFound := global.DB.First(&comment, commentID).RecordNotFound()
	if notFound {
		return gorm.ErrRecordNotFound
	}
	err = global.DB.Delete(&comment).Error
	return err
}
func QueryCommentByPostID(postID string) (comments []model.Comment, err error) {
	err = global.DB.Where("post_id = ?", postID).Find(&comments).Error
	return comments, err
}
