package router

import (
	v1 "gin-project/api/v1"

	"github.com/gin-gonic/gin"
)

func InitPostRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("api/v1/post")
	{
		UserRouter.POST("/create", v1.CreatePost)
		UserRouter.POST("/comment/create", v1.CreateComment)
		UserRouter.POST("/comment/list_all_comments", v1.ListAllComments)
	}
}
