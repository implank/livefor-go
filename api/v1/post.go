package v1

import (
	"gin-project/model"
	"gin-project/service"
	"gin-project/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreatePost doc
// @description  Create a post --note-- section in [0,3]
// @Tags         Post
// @Accept       json
// @Produce      json
// @Param        data  body      model.CreatePostData  true  "22"
// @Success      200   {string}  string                "{"success": true, "message": "发布成功"}"
// @Router       /post/create [post]
func CreatePost(c *gin.Context) {
	var data model.CreatePostData
	if err := c.ShouldBindJSON(&data); err != nil {
		panic(err)
	}
	user, notFound := service.QueryUserByUserID(data.UserID)
	if notFound {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户不存在",
		})
		return
	}
	if user.Bandate.After(time.Now()) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户禁言中",
		})
		return
	}
	post := model.Post{
		UserID:   data.UserID,
		Username: user.Username,
		Title:    data.Title,
		Content:  data.Content,
		Section:  data.Section,
	}
	service.CreatePost(&post)
	service.UpdateUserExp(user.UserID, 15)
	tags := data.Tags
	for _, tag := range tags {
		service.CreateTag(&tag, post.Section)
		postTag := model.PostTag{
			PostID: post.PostID,
			Name:   tag.Name,
		}
		service.CreatePostTag(&postTag)
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "发布成功",
	})
}

// func DeletePost(c *gin.Context) {
// 	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
// 	user, notFound := service.QueryUserByUserID(userID)
// 	if notFound || user.Username != "admin" {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": "用户不存在",
// 		})
// 		return
// 	}
// }

// GetPosts doc
// @description  Get posts with offset and length
// @Tags         Post
// @Accept       json
// @Produce      json
// @Param        data  body      model.GetPostsData  true  "22"
// @Success      200     {string}  string  "{"success": true, "message": "获取文章成功", "data": data}"
// @Router       /post/get [post]
func GetPosts(c *gin.Context) {
	var d model.GetPostsData
	if err := c.ShouldBindJSON(&d); err != nil {
		panic(err)
	}
	var order string
	if d.Order == "new" {
		order = "create_time desc"
	} else if d.Order == "top" {
		order = "'like' desc"
	} else if d.Order == "hot" {
		order = "views desc"
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "order参数错误",
		})
		return
	}
	posts, count := service.GetPosts(d.Offset, d.Length, d.Section, order, d.Tags)
	var data [](map[string]interface{})
	for _, post := range posts {
		tags, _ := service.QueryPostTags(post.PostID)
		_, notFound := service.QueryPostLike(post.PostID, d.UserID)
		user, _ := service.QueryUserByUserID(post.UserID)
		var like int
		if notFound {
			like = 0
		} else {
			like = 1
		}
		data = append(data, map[string]interface{}{
			"post":                post,
			"tags":                tags,
			"create_time_seconds": time.Since(post.CreateTime).Seconds(),
			"create_time_minutes": time.Since(post.CreateTime).Minutes(),
			"create_time_hours":   time.Since(post.CreateTime).Hours(),
			"create_time_days":    time.Since(post.CreateTime).Hours() / 24,
			"create_time_weeks":   time.Since(post.CreateTime).Hours() / 24 / 7,
			"like_status":         like,
			"user":                user,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "获取文章成功",
		"data":    data,
		"count":   count,
	})
}

// SearchPost doc
// @description	Search post
// @Tags         Post
// @Accept       json
// @Produce      json
// @Param        data  body      model.SearchPostsData  true  "22"
// @Success      200     {string}  string  "{"success": true, "message": "搜索成功", "data": data}"
// @Router       /post/search [post]
func SearchPosts(c *gin.Context) {
	var d model.SearchPostsData
	if err := c.ShouldBindJSON(&d); err != nil {
		panic(err)
	}
	var order string
	if d.Order == "new" {
		order = "create_time desc"
	} else if d.Order == "top" {
		order = "'like' desc"
	} else if d.Order == "hot" {
		order = "views desc"
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "order参数错误",
		})
		return
	}
	posts, count := service.SearchPosts(d.Fliters, d.Section, d.Offset, d.Length, order)
	var data [](map[string]interface{})
	for _, post := range posts {
		tags, _ := service.QueryPostTags(post.PostID)
		_, notFound := service.QueryPostLike(post.PostID, d.UserID)
		user, _ := service.QueryUserByUserID(post.UserID)
		var like int
		if notFound {
			like = 0
		} else {
			like = 1
		}
		data = append(data, map[string]interface{}{
			"post":                post,
			"tags":                tags,
			"create_time_seconds": time.Since(post.CreateTime).Seconds(),
			"create_time_minutes": time.Since(post.CreateTime).Minutes(),
			"create_time_hours":   time.Since(post.CreateTime).Hours(),
			"create_time_days":    time.Since(post.CreateTime).Hours() / 24,
			"create_time_weeks":   time.Since(post.CreateTime).Hours() / 24 / 7,
			"like_status":         like,
			"user":                user,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "获取文章成功",
		"data":    data,
		"count":   count,
	})
}

// GetUserPosts doc
// @description	Get user posts
// @Tags         Post
// @Param        user_id  formData  string  true  "user_id"
// @Param        offset  formData  string  true  "offset"
// @Param        length  formData  string  true  "length"
// @Success      200     {string}  string  "{"success": true, "message": "获取文章成功", "data": data}"
// @Router       /post/get_user_posts [post]
func GetUserPosts(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	off, _ := strconv.ParseUint(c.Request.FormValue("offset"), 0, 64)
	len, _ := strconv.ParseUint(c.Request.FormValue("length"), 0, 64)
	posts, count := service.GetUserPosts(userID, off, len)
	var data [](map[string]interface{})
	for _, post := range posts {
		tags, _ := service.QueryPostTags(post.PostID)
		data = append(data, map[string]interface{}{
			"post": post,
			"tags": tags,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "获取用户文章成功",
		"data":    data,
		"count":   count,
	})
}

// LikePost doc
// @description Like a post
// @Tags 				Post
// @Accept       json
// @Produce      json
// @Param        data  body      model.LikePostData  true  "LikePostData"
// @Success      200   {string}  string                 "{"success": true, "message": "点赞成功", "postlike": postlike}"
// @Router       /post/like [post]
func LikePost(c *gin.Context) {
	var data model.LikePostData
	if err := c.ShouldBindJSON(&data); err != nil {
		panic(err)
	}
	postLike, notFound := service.QueryPostLike(data.PostID, data.UserID)
	user, _ := service.QueryUserByUserID(data.UserID)
	post, _ := service.QueryPost(data.PostID)
	if !notFound {
		post.Like -= 1
		service.DeletePostLike(&postLike)
		service.UpdatePost(&post)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "取消点赞成功",
		})
		return
	}
	postLike = model.PostLike(data)
	post.Like += 1
	post.MaxLike = utils.Max(post.Like, post.MaxLike)
	notification := model.Notification{
		UserID:   post.UserID,
		Username: "@" + user.Username,
		PostID:   post.PostID,
		Msg:      "点赞了你的帖子",
		Content:  "\"" + post.Title + "\"",
	}
	if post.MaxLike == 10 || post.MaxLike == 50 || post.MaxLike == 150 {
		service.UpdateUserExp(post.UserID, 5)
		post.MaxLike += 1
	}
	if post.MaxLike == 500 || post.MaxLike == 1000 {
		service.UpdateUserExp(post.UserID, 10)
		post.MaxLike += 1
	}
	service.CreateNotification(&notification)
	service.CreatePostLike(&postLike)
	service.UpdatePost(&post)
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"message":  "点赞成功",
		"postLike": postLike,
	})
}

// CreateComment doc
// @description  Create a comment
// @Tags         Post
// @Param        user_id  formData  string  true  "user_id"
// @Param        post_id  formData  string  true  "post_id"
// @Param        content  formData  string  true  "content"
// @Success      200      {string}  string  "{"success": true, "message": "用户评论成功"}"
// @Router       /post/comment/create [post]
func CreateComment(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	user, _ := service.QueryUserByUserID(userID)
	postID, _ := strconv.ParseUint(c.Request.FormValue("post_id"), 0, 64)
	content := c.Request.FormValue("content")
	comment := model.Comment{
		UserID:      userID,
		Username:    user.Username,
		PostID:      postID,
		CommentTime: time.Now(),
		Content:     content,
	}
	err := service.CreateComment(&comment)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "评论失败",
		})
		return
	}
	post, _ := service.QueryPost(postID)
	post.Comment += 1
	if post.Comment == 10 || post.Comment == 30 || post.Comment == 60 ||
		post.Comment == 100 || post.Comment == 200 {
		service.UpdateUserExp(post.UserID, 5)
	}
	post.LastCommentTime = comment.CommentTime
	service.UpdatePost(&post)
	if len(content) > 15 {
		service.UpdateUserExp(userID, 5)
	}
	notification := model.Notification{
		UserID:   post.UserID,
		Username: "@" + user.Username,
		PostID:   post.PostID,
		Msg:      "评论了你的帖子",
		Content:  "\"" + post.Title + "\"",
	}
	service.CreateNotification(&notification)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "用户评论成功",
	})
}

// LikeComment doc
// @description  Like a comment
// @Tags         Post
// @Accept       json
// @Produce      json
// @Param        data  body      model.LikeCommentData  true  "LikeCommentData"
// @Success      200   {string}  string                 "{"success": true, "message": "点赞成功", "commentlike": commentlike}"
// @Router       /post/comment/like [post]
func LikeComment(c *gin.Context) {
	var data model.LikeCommentData
	var msg string
	if err := c.ShouldBindJSON(&data); err != nil {
		panic(err)
	}
	commentLike, notFound := service.QueryCommentLike(data.CommentID, data.UserID)
	comment, _ := service.QueryComment(data.CommentID)
	if data.LikeOrDislike {
		msg = "点赞"
	} else {
		msg = "踩"
	}
	if !notFound {
		if commentLike.LikeOrDislike {
			comment.Like -= 1
		} else {
			comment.Dislike -= 1
		}
		service.DeleteCommentLike(&commentLike)
		if commentLike.LikeOrDislike == data.LikeOrDislike {
			service.UpdateComment(&comment)
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "取消" + msg + "成功",
			})
			return
		}
	}
	commentLike = model.CommentLike(data)
	if commentLike.LikeOrDislike {
		comment.Like += 1
	} else {
		comment.Dislike += 1
	}
	comment.MaxLike = utils.Max(comment.Like, comment.MaxLike)
	if data.LikeOrDislike {
		user, _ := service.QueryUserByUserID(data.UserID)
		post, _ := service.QueryPost(comment.PostID)
		notification := model.Notification{
			UserID:   comment.UserID,
			Username: "@" + user.Username,
			PostID:   comment.PostID,
			Msg:      "点赞了你的评论",
			Content:  "\"" + post.Title + "\"",
		}
		service.CreateNotification(&notification)
		if comment.MaxLike == 10 || comment.MaxLike == 50 || comment.MaxLike == 150 {
			service.UpdateUserExp(comment.UserID, 5)
			comment.MaxLike += 1
		}
		if comment.MaxLike == 500 || comment.MaxLike == 1000 {
			service.UpdateUserExp(comment.UserID, 10)
			comment.MaxLike += 1
		}
	}
	service.CreateCommentLike(&commentLike)
	service.UpdateComment(&comment)
	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"message":     msg + "成功",
		"commentLike": commentLike,
	})
}

// GetPostComments doc
// @description  Get post comments
// @Tags         Post
// @Param				 user_id  formData  string  true  "user_id"
// @Param        post_id  formData  string  true  "post_id"
// @Param 			 offset formData  string  true  "offset"
// @Param 			 length formData  string  true  "length"
// @Success      200      {string}  string  "{"success": true, "message": "获取评论成功", "comments":comments}"
// @Router       /post/get_post_comments [post]
func GetPostComments(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	postID, _ := strconv.ParseUint(c.Request.FormValue("post_id"), 0, 64)
	off, _ := strconv.ParseUint(c.Request.FormValue("offset"), 0, 64)
	len, _ := strconv.ParseUint(c.Request.FormValue("length"), 0, 64)
	comments, count, err := service.GetPostComments(off, len, postID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "获取评论失败",
		})
		return
	}
	post, notFound := service.QueryPost(postID)
	if notFound {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "postID错误",
		})
		return
	}
	post.Views += 1
	if post.Views == 60 || post.Views == 200 || post.Views == 500 {
		service.UpdateUserExp(post.UserID, 5)
	}
	service.UpdateUserExp(userID, 5)
	service.UpdatePost(&post)
	var data [](map[string]interface{})
	for _, comment := range comments {
		var like int
		commentLike, notFound := service.QueryCommentLike(comment.CommentID, userID)
		user, _ := service.QueryUserByUserID(comment.UserID)
		if notFound {
			like = 0
		} else if commentLike.LikeOrDislike {
			like = 1
		} else {
			like = -1
		}
		data = append(data, map[string]interface{}{
			"comment":              comment,
			"comment_time_seconds": time.Since(comment.CommentTime).Seconds(),
			"comment_time_minutes": time.Since(comment.CommentTime).Minutes(),
			"comment_time_hours":   time.Since(comment.CommentTime).Hours(),
			"comment_time_days":    time.Since(comment.CommentTime).Hours() / 24,
			"comment_time_weeks":   time.Since(comment.CommentTime).Hours() / 24 / 7,
			"like_status":          like,
			"user":                 user,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "获取评论成功",
		"data":    data,
		"count":   count,
	})
}

// AddPostTag doc
// @description  Add post a tag
// @Tags         Post
// @Param        post_id  formData  string  true  "post_id"
// @Param        name  formData  string  true  "name"
// @Success      200      {string}  string  "{"success": true, "message": "添加标签成功"}"
// @Router       /post/add_post_tag [post]
func AddPostTag(c *gin.Context) {
	post_id, _ := strconv.ParseUint(c.Request.FormValue("post_id"), 0, 64)
	name := c.Request.FormValue("name")
	post, notFound := service.QueryPost(post_id)
	if notFound {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "postID错误",
		})
		return
	}
	err := service.CreateTag(&model.Tag{Name: name}, post.Section)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "添加标签失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "添加标签成功",
	})
}

// GetPostTags doc
// @description  Get post tags
// @Tags         Post
// @Param        post_id  formData  string  true  "post_id"
// @Success      200      {string}  string  "{"success": true, "message": "获取标签成功", "post_tags":postTags}"
// @Router       /post/get_post_tags [post]
func GetPostTags(c *gin.Context) {
	postID, _ := strconv.ParseUint(c.Request.FormValue("post_id"), 0, 64)
	postTags, _ := service.QueryPostTags(postID)
	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"message":   "获取标签成功",
		"post_tags": postTags,
	})
}

// GetSectionTags doc
// @description  Get section tags
// @Tags         Post
// @Param				 section  query  string  true  "section"
// @Success      200  {string}  string  "{"success": true, "message": "获取标签成功", "tags":tags}"
// @Router       /post/get_section_tags [get]
func GetSectionTags(c *gin.Context) {
	section, _ := strconv.ParseUint(c.Query("section"), 0, 64)
	sectionTags, _ := service.QuerySectionTags(section)
	var tags []model.Tag
	for _, tag := range sectionTags {
		tags = append(tags, model.Tag{Name: tag.Name})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "获取标签成功",
		"data":    tags,
	})
}
