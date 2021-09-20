package routes

import (
	"github.com/labstack/echo/v4"
	"simple-users-mvc/controllers"
)
func NewRoute() *echo.Echo {
	e := echo.New()
	e.GET("/", controllers.GreetMessage)
	e.GET("/user", controllers.GetAllUsersController)
	e.POST("/user", controllers.PostUsersController)
	e.GET("/user/:id", controllers.GetUserByIDController)
	e.PUT("/user/:id", controllers.UpdateUserController)
	e.DELETE("/user/:id", controllers.DeleteUserController)
	
	return e
}
