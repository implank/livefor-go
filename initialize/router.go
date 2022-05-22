package initialize

import (
	"gin-project/middleware"
	"gin-project/router"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	Group := r.Group("")
	{
		router.InitUserRouter(Group) // 注册用户路由
		router.InitPostRouter(Group)
		router.InitPortalRouter(Group)
	}
	return r
}
