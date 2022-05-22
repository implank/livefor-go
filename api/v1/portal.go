package v1

import (
	"gin-project/model"
	"gin-project/service"
	"net/http"

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
