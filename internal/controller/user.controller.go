package controller

import (
	"go-ecomerce-backend-api/m/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// uc user controller
// us user service

// controller -> service -> repo -> models -> dbs
func (uc *UserController) GetUserByID(c *gin.Context) {
	// if err != nil {
	// 	return response.ErrorResponse(c, 20003, "Email is invalid")
	// }
	// return response.SuccessResponse(c, 20001, []string{"user1", "user2"})
}