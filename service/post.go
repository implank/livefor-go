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
func CreatePostTag(postTag *model.PostTag) (err error) {
	if err = global.DB.Create(postTag).Error; err != nil {
		return err
	}
	return
}
func CreateTag(tag *model.Tag) (err error) {
	if err = global.DB.Create(tag).Error; err != nil {
		return err
	}
	return
}
func QueryTag(name string) (tag *model.Tag, notFound bool) {
	notFound = global.DB.First(tag, name).RecordNotFound()
	return tag, notFound
}
func QueryPostTags(postID uint64) (postTags []model.PostTag, err error) {
	err = global.DB.Where("post_id = ?", postID).Find(&postTags).Error
	return postTags, err
}
func QueryAllTags() (tags []model.Tag, err error) {
	err = global.DB.Find(&tags).Error
	return tags, err
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
func QueryPostComments(postID uint64) (comments []model.Comment, err error) {
	err = global.DB.Where("post_id = ?", postID).Find(&comments).Error
	return comments, err
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
func CreateCommentLike(commentLike *model.CommentLike) (err error) {
	if err = global.DB.Create(commentLike).Error; err != nil {
		return err
	}
	var comment model.Comment
	global.DB.First(&comment, commentLike.CommentID)
	if commentLike.LikeOrDislike {
		comment.Like += 1
	} else {
		comment.Dislike += 1
	}
	global.DB.Save(&comment)
	return
}
func DeleteCommentLike(commentLike *model.CommentLike) (err error) {
	var comment model.Comment
	global.DB.First(&comment, commentLike.CommentID)
	if commentLike.LikeOrDislike {
		comment.Like -= 1
	} else {
		comment.Dislike -= 1
	}
	global.DB.Save(&comment)
	global.DB.Delete(&commentLike)
	return
}
func QueryCommentLike(commentID uint64, userID uint64) (commentLike model.CommentLike, notFound bool) {
	notFound = global.DB.First(&commentLike, "comment_id = ? and user_id = ?", commentID, userID).RecordNotFound()
	return commentLike, notFound
}
