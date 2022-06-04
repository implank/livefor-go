package router

import (
	v1 "gin-project/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("/user")
	{
		UserRouter.POST("/register", v1.Register)
		UserRouter.POST("/login", v1.Login)
		UserRouter.POST("/info", v1.ShowUserInfo)
		UserRouter.POST("/update_password", v1.UpdatePassword)
		UserRouter.POST("/update_info", v1.UpdateInfo)
		UserRouter.POST("/upload_avatar", v1.UploadAvatar)
		UserRouter.POST("/test", v1.Test)
	}
	PostRouter := Router.Group("/post")
	{
		PostRouter.POST("/create", v1.CreatePost)
		PostRouter.POST("/get", v1.GetPosts)
		PostRouter.POST("/like", v1.LikePost)
		PostRouter.POST("/comment/create", v1.CreateComment)
		PostRouter.POST("/comment/like", v1.LikeComment)
		PostRouter.POST("/get_post_comments", v1.GetPostComments)
		PostRouter.POST("/get_post_tags", v1.GetPostTags)
		PostRouter.POST("/get_all_tags", v1.GetAllTags)
	}
	PortalRouter := Router.Group("/portal")
	{
		PortalRouter.POST("/save_greenbirds", v1.SaveGreenbird)
		PortalRouter.POST("/get_greenbirds", v1.GetGreenbird)
		PortalRouter.POST("/ban_user", v1.BanUser)
		PortalRouter.POST("/get_banned_users", v1.GetBannedUsers)
		PortalRouter.POST("/upload_file", v1.UploadFile)
	}
}
