package routes

import (
	controllers "split_it_backend/controllers/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController controllers.UserController
	JWTConfig	echojwt.Config
}

func(controllerList *ControllerList) RouteRegister(e *echo.Echo){
	users := e.Group("/user")
	users.POST("/register",controllerList.UserController.Register)
	users.POST("/login",controllerList.UserController.Login)

}