package router

import (
	v1 "gin-project/api/v1"

	"github.com/gin-gonic/gin"
)

func InitPostRouter(Router *gin.RouterGroup) {
	PostRouter := Router.Group("api/v1/post")
	{
		PostRouter.POST("/create", v1.CreatePost)
		PostRouter.POST("/get", v1.GetPosts)
		PostRouter.POST("/comment/create", v1.CreateComment)
		PostRouter.POST("/comment/like", v1.LikeComment)
		PostRouter.POST("/get_post_comments", v1.GetPostComments)
		PostRouter.POST("/get_post_tags", v1.GetPostTags)
		PostRouter.POST("/get_all_tags", v1.GetAllTags)
	}
}
