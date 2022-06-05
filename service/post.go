package service

import (
	"gin-project/global"
	"gin-project/model"
	"strings"

	"github.com/jinzhu/gorm"
)

func CreatePost(post *model.Post) (err error) {
	if err = global.DB.Create(post).Error; err != nil {
		return err
	}
	return
}
func QueryPost(postID uint64) (post model.Post, notFound bool) {
	notFound = global.DB.First(&post, postID).RecordNotFound()
	return post, notFound
}
func DeletePost(postID uint64) (err error) {
	var post model.Post
	notFound := global.DB.First(&post, postID).RecordNotFound()
	if notFound {
		return gorm.ErrRecordNotFound
	}
	err = global.DB.Delete(&post).Error
	return err
}
func UpdatePost(post *model.Post) (err error) {
	err = global.DB.Save(post).Error
	return
}
func GetPosts(off uint64, leng uint64, section uint64, order string, tags []model.Tag) (
	post []model.Post, count uint64) {
	if len(tags) == 0 {
		global.DB.Order(order).Where("section = ?", section).
			Limit(leng).Offset(off).Find(&post).
			Limit(-1).Offset(-1).Count(&count)
	} else {
		var str string
		var id []string
		for _, tag := range tags {
			id = append(id, "'"+tag.Name+"'")
		}
		str = strings.Join(id, ",")
		global.DB.
			Raw(
				"select *	from posts where exists("+
					" select * from post_tags "+
					" where post_tags.post_id=posts.post_id and post_tags.name in ("+str+"))"+
					" and section = ?", section).
			Find(&post)
		count = (uint64)(len(post))
		global.DB.Order(order).Limit(leng).Offset(off).
			Raw(
				"select *	from posts where exists( "+
					" select * from post_tags "+
					" where post_tags.post_id=posts.post_id and post_tags.name in ("+str+"))"+
					" and section = ?", section).
			Find(&post)
	}
	return post, count
}
func GetUserPosts(
	userID uint64, off uint64, len uint64) (
	post []model.Post, count uint64) {
	global.DB.Order("create_time desc").Where("user_id = ?", userID).
		Limit(len).Offset(off).Find(&post).
		Limit(-1).Offset(-1).Count(&count)
	return post, count
}
func CreatePostLike(postLike *model.PostLike) (err error) {
	if err = global.DB.Create(postLike).Error; err != nil {
		return err
	}
	var post model.Post
	global.DB.First(&post, postLike.PostID)
	if postLike.LikeOrDislike {
		post.Like += 1
	} else {
		post.Dislike += 1
	}
	global.DB.Save(&post)
	return
}
func DeletePostLike(postLike *model.PostLike) (err error) {
	var post model.Post
	global.DB.First(&post, postLike.PostID)
	if postLike.LikeOrDislike {
		post.Like -= 1
	} else {
		post.Dislike -= 1
	}
	global.DB.Save(&post)
	global.DB.Delete(&postLike)
	return
}
func QueryPostLike(
	postID uint64, userID uint64) (
	postLike model.PostLike, notFound bool) {
	notFound = global.DB.First(&postLike,
		"post_id = ? and user_id = ?", postID, userID).RecordNotFound()
	return postLike, notFound
}
func CreatePostTag(postTag *model.PostTag) (err error) {
	if err = global.DB.Create(postTag).Error; err != nil {
		return err
	}
	return
}
func CreateTag(tag *model.Tag) (err error) {
	_, notFound := QueryTag(tag.Name)
	if !notFound {
		return
	}
	if err = global.DB.Create(tag).Error; err != nil {
		return err
	}
	return
}
func QueryTag(name string) (tag model.Tag, notFound bool) {
	notFound = global.DB.Where("name = ?", name).First(&tag).RecordNotFound()
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
func GetPostComments(
	off uint64, len uint64, postID uint64) (
	comments []model.Comment, count uint64, err error) {
	err = global.DB.Order("comment_time").Where("post_id = ?", postID).
		Limit(len).Offset(off).Find(&comments).
		Limit(-1).Offset(-1).Count(&count).Error
	return comments, count, err
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
func QueryCommentLike(
	commentID uint64, userID uint64) (
	commentLike model.CommentLike, notFound bool) {
	notFound = global.DB.First(&commentLike,
		"comment_id = ? and user_id = ?", commentID, userID).RecordNotFound()
	return commentLike, notFound
}
