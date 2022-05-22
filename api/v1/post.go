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
// @description  Create a post
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
	post := model.Post{
		UserID:  data.UserID,
		Title:   data.Title,
		Content: data.Content,
	}
	service.CreatePost(&post)
	tags := data.Tags
	for _, tag := range tags {
		service.CreateTag(&tag)
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
	commentLike, notFoundCommentLike := service.QueryCommentLike(data.CommentID, data.UserID)
	if !notFoundCommentLike {
		service.DeleteCommentLike(&commentLike)
	}
	commentLike = model.CommentLike{
		CommentID:     data.CommentID,
		UserID:        data.UserID,
		LikeOrDislike: data.LikeOrDislike,
	}
	service.CreateCommentLike(&commentLike)
	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"message":     "LikeOrDislike成功",
		"commentLike": commentLike,
	})
}

// GetPostComments doc
// @description  Get post comments
// @Tags         Post
// @Param        post_id  formData  string  true  "post_id"
// @Success      200      {string}  string  "{"success": true, "message": "获取评论成功", "comments":comments}"
// @Router       /post/get_post_comments [post]
func GetPostComments(c *gin.Context) {
	postID, _ := strconv.ParseUint(c.Request.FormValue("post_id"), 0, 64)
	comments, err := service.QueryPostComments(postID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "获取评论失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"message":  "获取评论成功",
		"comments": comments,
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

// GetAllTags doc
// @description  Get all tags
// @Tags         Post
// @Success      200  {string}  string  "{"success": true, "message": "获取标签成功", "tags":tags}"
// @Router       /post/get_all_tags [post]
func GetAllTags(c *gin.Context) {
	tags, err := service.QueryAllTags()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "获取标签失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "获取标签成功",
		"data":    tags,
	})
}
