package routes

import (
	"gin-gorm-postgres/controllers"
	"gin-gorm-postgres/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewUserRouteController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(c *gin.RouterGroup) {
	r := c.Group("/users")
	r.GET("/me", middleware.DeserializeUser(), uc.userController.GetMe)
}
