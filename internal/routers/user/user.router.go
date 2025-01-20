package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// Public router
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.GET("/register")
		userRouterPublic.POST("/otp")
	}
	// Private router
	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.GET("/get_info")
	}
}