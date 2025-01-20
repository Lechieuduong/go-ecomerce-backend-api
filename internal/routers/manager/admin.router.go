package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (ur *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// Public router
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}
	// Private router
	adminRouterPrivate := Router.Group("/admin/user")
	{
		adminRouterPrivate.GET("/active_user")
	}
}