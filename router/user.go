package router

import (
	v1 "gin-project/api/v1"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("api/v1/user")
	{
		UserRouter.POST("/register", v1.Register)
		UserRouter.POST("/login", v1.Login)
		UserRouter.POST("/info", v1.ShowUserInfo)
		UserRouter.POST("/update_password", v1.UpdatePassword)
		UserRouter.POST("/update_info", v1.UpdateInfo)
		UserRouter.POST("/upload_avatar", v1.UploadAvatar)
		UserRouter.POST("/test", v1.Test)
	}
}
