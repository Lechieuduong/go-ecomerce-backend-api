package routers

import (
	"go-ecomerce-backend-api/m/internal/routers/manager"
	"go-ecomerce-backend-api/m/internal/routers/user"
)

type RouterGroup struct {
	User user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)