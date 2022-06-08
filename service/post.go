package service

import (
	"gin-project/global"
	"gin-project/model"
	"gin-project/utils"
	"strings"

	"github.com/jinzhu/gorm"
)

func CreatePost(post *model.Post) (err error) {
	err = global.DB.Create(post).Error
	return err
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
func SearchPosts(
	fliters []string, section uint64, off uint64, lim uint64, order string) (
	posts []model.Post, count uint64) {
	var tmp []model.Post
	if section == 1926 {
		global.DB.Order(order).Find(&tmp)
	} else {
		global.DB.Order(order).Where("section = ?", section).Find(&tmp)
	}
	for _, post := range tmp {
		for _, fliter := range fliters {
			if strings.Contains(post.Title, fliter) ||
				strings.Contains(post.Content, fliter) {
				posts = append(posts, post)
			}
		}
	}
	count = (uint64)(len(posts))
	lim = utils.Min(count-off, lim)
	posts = posts[off : off+lim]
	return posts, count
}
func GetHotPosts() (posts []model.Post) {
	global.DB.Order("views desc").Limit(4).Find(&posts)
	return posts
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
	err = global.DB.Create(postTag).Error
	return err
}
func CreateTag(tag *model.Tag, section uint64) (err error) {
	sectionTag, notFound := QuerySectionTag(tag.Name, section)
	if !notFound {
		sectionTag.Infers += 1
		UpdateTag(&sectionTag)
		return
	}
	sectionTag = model.SectionTag{Name: tag.Name, Section: section}
	err = global.DB.Create(&sectionTag).Error
	return err
}
func UpdateTag(tag *model.SectionTag) {
	global.DB.Save(tag)
}
func QuerySectionTag(name string, section uint64) (tag model.SectionTag, notFound bool) {
	notFound = global.DB.Where("name = ? and section = ?", name, section).
		First(&tag).RecordNotFound()
	return tag, notFound
}
func QueryPostTags(postID uint64) (postTags []model.PostTag, notFound bool) {
	notFound = global.DB.Where("post_id = ?", postID).
		Find(&postTags).RecordNotFound()
	return postTags, notFound
}
func QuerySectionTags(section uint64) (tags []model.SectionTag, notFound bool) {
	notFound = global.DB.Order("Infers desc").Where("section = ?", section).
		Find(&tags).RecordNotFound()
	return tags, notFound
}
func CreateComment(comment *model.Comment) (err error) {
	err = global.DB.Create(comment).Error
	return err
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
