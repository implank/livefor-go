package v1

import (
	"fmt"
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

// Register doc
// @Description  Register
// @Tags         User
// @Param        username   formData  string  true  "username"
// @Param        password1     formData  string  true      "password1"
// @Param        password2     formData  string  true      "password2"
// @Param        email      formData  string  true  "email"
// @Success      200        {string}  string  "{"status": true, "message": "注册成功"}"
// @Router       /user/register [post]
func Register(c *gin.Context) {
	username := c.Request.FormValue("username")
	password1 := c.Request.FormValue("password1")
	password2 := c.Request.FormValue("password2")
	email := c.Request.FormValue("email")
	if password1 != password2 {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "两次输入的密码不一致",
		})
		return
	}
	_, notFoundUsername := service.QueryUserByUsername(username)
	_, notFoundEmail := service.QueryUserByEmail(email)
	if notFoundUsername && notFoundEmail {
		user := model.User{
			Username: username,
			Password: password1,
			Email:    email,
		}
		service.CreateUser(&user)
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "注册成功",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  false,
		"message": "用户名/邮箱已存在",
	})
}

// Login doc
// @Description  Login
// @Tags         User
// @Param        username  formData  string  true      "username"
// @Param        password  formData  string  true  "password"
// @Success      200       {string}  string  "{"status": true, "message": "登录成功","data": user}"
// @Router       /user/login [post]
func Login(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	user, notFound := service.QueryUserByUsername(username)
	if notFound {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "用户名不存在",
		})
		return
	}
	if user.Password != password {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "密码错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "登录成功",
		"data":    user,
	})
}

// ShowUserInfo doc
// @Description  ShowUserInfo
// @Tags         User
// @Param        user_id       formData  string  true      "user_id"
// @Success      200      {string}  string  "{"status": true, "message": "查询成功", "data": user}"
// @Router       /user/info [post]
func ShowUserInfo(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	user, notFound := service.QueryUserByUserID(userID)
	if notFound {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "用户名不存在",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "查询成功",
		"data":    user,
	})
}

// UpdatePassword doc
// @Description  UpdatePassword
// @Tags         User
// @Param        user_id   formData  string  true      "user_id"
// @Param        old_password  formData  string  true      "old_password"
// @Param        password1  formData  string  true  "password1"
// @Param        password2  formData  string  true  "password2"
// @Success                              200     {string}  string  "{"status": true, "message": "修改成功"}"
// @Router       /user/update_password [post]
func UpdatePassword(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	oldPassword := c.Request.FormValue("old_password")
	password1 := c.Request.FormValue("password1")
	password2 := c.Request.FormValue("password2")
	if oldPassword == password1 {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "新密码不能与旧密码相同",
		})
		return
	}
	if password1 != password2 {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "两次输入的密码不一致",
		})
		return
	}
	user, notFound := service.QueryUserByUserID(userID)
	if notFound {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "用户不存在",
		})
		return
	}
	if user.Password != oldPassword {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "旧密码错误",
		})
		return
	}
	user.Password = password1
	service.UpdateUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "修改成功",
	})
}

// UpdateInfo doc
// @Description  UpdateInfo
// @Tags         User
// @Param        user_id  formData  string  true  "user_id"
// @Param        username  formData  string  true  "username"
// @Param        email     formData  string  true      "email"
// @Param        sex       formData  string  false      "sex"
// @Param        age       formData  string  false      "age"
// @Success      200     {string}  string  "{"status": true, "message": "修改成功"}"
// @Router       /user/update_info [post]
func UpdateInfo(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	username := c.Request.FormValue("username")
	email := c.Request.FormValue("email")
	sex := c.Request.FormValue("sex")
	age, _ := strconv.ParseUint(c.Request.FormValue("age"), 0, 64)
	user, notFound := service.QueryUserByUserID(userID)
	if notFound {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "用户不存在",
		})
		return
	}
	puser, notFound := service.QueryUserByUsername(username)
	if !notFound && puser.UserID != user.UserID {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "用户名已存在",
		})
		return
	}
	puser, notFound = service.QueryUserByEmail(email)
	if !notFound && puser.UserID != user.UserID {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "邮箱已存在",
		})
		return
	}
	user.Username = username
	user.Email = email
	user.Sex = sex
	user.Age = age
	service.UpdateUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "修改成功",
	})
}

// UploadAvatar doc
// @Description  upload user avatar
// @Tags         User
// @Param 			 avatar  formData  file true "avatar"
// @Param        user_id  formData  string  true  "user_id"
// @Success      200     {string}  string  "{"status": true, "message": "上传成功", "avatar_url": url}"
// @Router       /user/upload_avatar [post]
func UploadAvatar(c *gin.Context) {
	_, header, _ := c.Request.FormFile("avatar")
	userid, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	user, _ := service.QueryUserByUserID(userid)
	raw := fmt.Sprintf("%d", userid) + time.Now().String() + header.Filename
	md5 := utils.GetMd5(raw)
	suffix := strings.Split(header.Filename, ".")[1]
	saveDir := "./media/avatars"
	saveName := md5 + "." + suffix
	savePath := path.Join(saveDir, saveName)
	c.SaveUploadedFile(header, savePath)
	url := "http://43.138.77.133:81/media/avatars" + saveName
	user.AvatarUrl = url
	service.UpdateUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"message":    "上传成功",
		"avatar_url": url,
	})
}
