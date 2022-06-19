package routes

import (
	"api/blueprints/users/controllers"

	"github.com/labstack/echo"
)

func New(e *echo.Echo) {

	e.POST("/auth/register", controllers.UserRegister)
	e.GET("/users", controllers.GetAllUsers)

}
