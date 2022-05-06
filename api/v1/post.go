package v1

import (
	"gin-project/model"
	"gin-project/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//CreatePost doc
// @description Create a post
// @Tags Post
// @Param user_id 	formData  string  true  "user_id"
// @Param title formData string true "title"
// @Param content formData string true "content"
// @Success 200 {string} string "{"success": true, "message": "用户发布成功"}"
// @Router /post/create [post]
func CreatePost(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	title := c.Request.FormValue("title")
	content := c.Request.FormValue("content")
	post := model.Post{
		UserID:  userID,
		Title:   title,
		Content: content,
	}
	err := service.CreatePost(&post)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "发布失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "发布成功",
	})
}

//CreateComment doc
// @description Create a comment
// @Tags Post
// @Param user_id formData string true "user_id"
// @Param post_id formData string true "post_id"
// @Param content formData string true "content"
// @Success 200 {string} string "{"success": true, "message": "用户评论成功"}"
// @Router /post/comment/create [post]
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

// func LikeComment(c *gin.Context) {
// 	commentID, _ := strconv.ParseUint(c.Request.FormValue("comment_id"), 0, 64)
// 	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
// 	_, commentNotFound := service.QueryComment(commentID)
// 	if commentNotFound {
// 		c.JSON(http.StatusOK, gin.H{
// 			"success": false,
// 			"message": "评论不存在",
// 		})
// 		return
// 	}
// 	commentLike := model.CommentLike{
// 		CommentID:       commentID,
// 		UserID:          userID,
// 		like_or_dislike: true,
// 	}
// 	err:=service.CreateCommentLike()
// }

//ListAllComment doc
// @description List all comments
// @Tags Post
// @Param post_id formData string true "post_id"
// @Success 200 {string} string "{"success": true, "message": "获取评论成功", "data":comments}"
// @Router /post/comment/list_all_comments [post]
func ListAllComments(c *gin.Context) {
	postID := c.Request.FormValue("post_id")
	comments, err := service.QueryCommentByPostID(postID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "获取评论失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "获取评论成功",
		"data":    comments,
	})
}
