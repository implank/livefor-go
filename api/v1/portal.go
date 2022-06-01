package v1

import (
	"gin-project/model"
	"gin-project/service"
	"net/http"
	"strconv"
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
