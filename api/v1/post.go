package v1

import (
	"gin-project/model"
	"gin-project/service"
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
	println(d.Order)
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
		data = append(data, map[string]interface{}{
			"post":                post,
			"tags":                tags,
			"create_time_seconds": time.Since(post.CreateTime).Seconds(),
			"create_time_minutes": time.Since(post.CreateTime).Minutes(),
			"create_time_hours":   time.Since(post.CreateTime).Hours(),
			"create_time_days":    time.Since(post.CreateTime).Hours() / 24,
			"create_time_weeks":   time.Since(post.CreateTime).Hours() / 24 / 7,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "获取文章成功",
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
	if !notFound {
		service.DeletePostLike(&postLike)
		if postLike.LikeOrDislike == data.LikeOrDislike {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "取消点赞成功",
			})
			return
		}
	}
	postLike = model.PostLike(data)
	service.CreatePostLike(&postLike)
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
	post, _ := service.QueryPost(postID)
	post.Comment += 1
	post.LastCommentTime = comment.CommentTime
	service.UpdatePost(&post)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "评论失败",
		})
		return
	}
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
	if err := c.ShouldBindJSON(&data); err != nil {
		panic(err)
	}
	var msg string
	if data.LikeOrDislike {
		msg = "点赞"
	} else {
		msg = "踩"
	}
	commentLike, notFound := service.QueryCommentLike(data.CommentID, data.UserID)
	if !notFound {
		service.DeleteCommentLike(&commentLike)
		if commentLike.LikeOrDislike == data.LikeOrDislike {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "取消" + msg + "成功",
			})
			return
		}
	}
	commentLike = model.CommentLike(data)
	service.CreateCommentLike(&commentLike)
	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"message":     msg + "成功",
		"commentLike": commentLike,
	})
}

// GetPostComments doc
// @description  Get post comments
// @Tags         Post
// @Param        post_id  formData  string  true  "post_id"
// @Param 			 offset formData  string  true  "offset"
// @Param 			 length formData  string  true  "length"
// @Success      200      {string}  string  "{"success": true, "message": "获取评论成功", "comments":comments}"
// @Router       /post/get_post_comments [post]
func GetPostComments(c *gin.Context) {
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
	service.UpdatePost(&post)
	var data [](map[string]interface{})
	for _, comment := range comments {
		data = append(data, map[string]interface{}{
			"comment":              comment,
			"comment_time_seconds": time.Since(comment.CommentTime).Seconds(),
			"comment_time_minutes": time.Since(comment.CommentTime).Minutes(),
			"comment_time_hours":   time.Since(comment.CommentTime).Hours(),
			"comment_time_days":    time.Since(comment.CommentTime).Hours() / 24,
			"comment_time_weeks":   time.Since(comment.CommentTime).Hours() / 24 / 7,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "获取评论成功",
		"data":    data,
		"count":   count,
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
	postTags, err := service.QueryPostTags(postID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "获取标签失败",
		})
		return
	}
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
	println(section)
	sectionTags := service.QuerySectionTags(section)
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
