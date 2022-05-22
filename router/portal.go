package router

import (
	v1 "gin-project/api/v1"

	"github.com/gin-gonic/gin"
)

func InitPortalRouter(Router *gin.RouterGroup) {
	PortalRouter := Router.Group("api/v1/portal")
	{
		PortalRouter.POST("/save_greenbirds", v1.SaveGreenbird)
		PortalRouter.POST("/get_greenbirds", v1.GetGreenbird)
	}
}
