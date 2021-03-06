package v1

import (
	"gin-project/model"
	"gin-project/service"
	"gin-project/utils"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// SaveGreenbird doc
// @Description  SaveGreenbird
// @Tags         Portal
// @Accept       json
// @Produce      json
// @Param        data  body      model.GreenbirdData  true  "新手上路信息"
// @Success      200   {string}  string               "{"status": true, "message": "保存成功"}"
// @Router       /portal/save_greenbirds [post]
func SaveGreenbird(c *gin.Context) {
	var data model.GreenbirdData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}
	err := service.SaveGreenbird(data.Greenbirds)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "保存成功",
	})
}

// GetGreenbirds doc
// @Description  GetGreenbirds
// @Tags         Portal
// @Success      200  {string}  string  "{"status": true, "message": "获取成功", "data": data}"
// @Router       /portal/get_greenbirds [post]
func GetGreenbird(c *gin.Context) {
	greenbirds, err := service.GetGreenbirds()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "获取成功",
		"data":    greenbirds,
	})
}

// BanUser doc
// @Description  BanUser
// @Tags         Portal
// @Param        user_id  query     int     true  "用户ID"
// @Success      200      {string}  string  "{"status": true, "message": "禁言成功"}"
// @Router       /portal/ban_user [post]
func BanUser(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	user, notFound := service.QueryUserByUserID(userID)
	if notFound {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户不存在",
		})
		return
	}
	user.Bandate = time.Now().Add(3 * time.Hour)
	service.UpdateUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "禁言成功",
	})
}

// UnbanUser doc
// @Description  UnbanUser
// @Tags         Portal
// @Param        user_id  query     int     true  "用户ID"
// @Success      200      {string}  string  "{"status": true, "message": "解禁成功"}"
// @Router       /portal/unban_user [post]
func UnbanUser(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	user, notFound := service.QueryUserByUserID(userID)
	if notFound {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户不存在",
		})
		return
	}
	user.Bandate = time.Now().Add(-1 * time.Hour)
	service.UpdateUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "解禁成功",
	})
}

// GetBannedUsers doc
// @Description  GetBannedUsers
// @Tags         Portal
// @Success      200  {string}  string  "{"status": true, "message": "获取成功", "users": users}"
// @Router       /portal/get_banned_users [post]
func GetBannedUsers(c *gin.Context) {
	users, _ := service.GetBannedUsers()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "获取成功",
		"users":   users,
	})
}

// UploadImage doc
// @Description  UploadImage
// @Tags         Portal
// @Param        image  formData   file     true  "图片"
// @Param				 user_id  formData  string  true      "user_id"
// @Success      200      {string}  string  "{"status": true, "message": "上传成功", "url": url}"
// @Router       /portal/upload_file [post]
func UploadFile(c *gin.Context) {
	_, header, _ := c.Request.FormFile("image")
	userid := c.Request.FormValue("user_id")
	raw := userid + time.Now().String() + header.Filename
	md5 := utils.GetMd5(raw)
	suffix := strings.Split(header.Filename, ".")[1]
	saveDir := "./media/images"
	saveName := md5 + "." + suffix
	savePath := path.Join(saveDir, saveName)
	c.SaveUploadedFile(header, savePath)
	url := "http://43.138.77.133:81/media/images/" + saveName
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "上传成功",
		"url":     url,
	})
}

// GetNotifications doc
// @description  Get user notifications
// @Tags         Portal
// @Param        user_id  formData     string     true  "user_id"
// @Param 			 offset 	formData  string  true  "offset"
// @Param 			 length 	formData  string  true  "length"
// @Param 			 type 	formData  string  true  "type"
// @Success      200      {string}  string  "{"status": true, "message": "获取成功", "data": notifications}"
// @Router       /portal/get_notifications [post]
func GetNotifications(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	off, _ := strconv.ParseUint(c.Request.FormValue("offset"), 0, 64)
	len, _ := strconv.ParseUint(c.Request.FormValue("length"), 0, 64)
	t, _ := strconv.ParseUint(c.Request.FormValue("type"), 0, 64)
	var notifications []model.Notification
	var count uint64
	if t == 0 {
		notifications, count = service.GetLikeNotification(userID, off, len)
	} else {
		notifications, count = service.GetCommentNotification(userID, off, len)
	}
	var data [](map[string]interface{})
	for _, notification := range notifications {
		user, _ := service.QueryUserByUserID(notification.UserID)
		data = append(data, map[string]interface{}{
			"notification": notification,
			"user_avatar":  user.AvatarUrl,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "获取成功",
		"data":    data,
		"count":   count,
	})
}

// GetUserMessage doc
// @description  Get user system message
// @Tags         Portal
// @Param        user_id  formData     string     true  "user_id"
// @Param 			 offset 	formData  string  true  "offset"
// @Param 			 length 	formData  string  true  "length"
// @Success      200      {string}  string  "{"status": true, "message": "获取成功", "data": SysMessages}"
// @Router       /portal/get_user_message [post]
func GetSysMessage(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	off, _ := strconv.ParseUint(c.Request.FormValue("offset"), 0, 64)
	len, _ := strconv.ParseUint(c.Request.FormValue("length"), 0, 64)
	SysMessages, count := service.GetSysMessages(userID, off, len)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "获取成功",
		"data":    SysMessages,
		"count":   count,
	})
}

// GetHotPosts doc
// @Description get posts which get the most highest views
// @Tags Portal
// @Success 200 {string} string "{"status": true, "message": "获取成功", "data": data}"
// @Router /portal/get_hot_posts [post]
func GetHotPosts(c *gin.Context) {
	posts := service.GetHotPosts()
	var data [](map[string]interface{})
	for _, post := range posts {
		data = append(data, map[string]interface{}{
			"post":                post,
			"create_time_seconds": time.Since(post.CreateTime).Seconds(),
			"create_time_minutes": time.Since(post.CreateTime).Minutes(),
			"create_time_hours":   time.Since(post.CreateTime).Hours(),
			"create_time_days":    time.Since(post.CreateTime).Hours() / 24,
			"create_time_weeks":   time.Since(post.CreateTime).Hours() / 24 / 7,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "获取成功",
		"data":    data,
	})
}

// GetGreenStatus doc
// @Description get user green status
// @Tags Portal
// @Param user_id formData string true "user_id"
// @Success 200 {string} string "{"status": true, "message": "获取成功", "data": greenStatus}"
// @Router /portal/get_green [post]
func GetGreen(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	t5 := service.GetSysMessageCount(userID, 5, "")
	t6 := service.GetSysMessageCount(userID, 6, "")
	c.JSON(http.StatusOK, gin.H{
		"success":          true,
		"completeuserinfo": t5,
		"uploaduseravatar": t6,
	})
}

// CheckNoob doc
// @Description check user is noob
// @Tags Portal
// @Param user_id formData string true "user_id"
// @Success 200 {string} string "{"status": true, "message": ""}"
// @Router /portal/check_noob [post]
func CheckNoob(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	user, _ := service.QueryUserByUserID(userID)
	user.Isnoob = false
	service.UpdateUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
}
