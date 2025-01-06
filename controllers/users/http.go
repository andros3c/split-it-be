package controllers

import (
	"net/http"
	"split_it_backend/businesses/users"
	"split_it_backend/controllers"

	"split_it_backend/controllers/users/request"
	"split_it_backend/controllers/users/response"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	usecase users.UseCase
}

func NewUserController(uc users.UseCase)*UserController{
	return &UserController{
		usecase: uc ,
	}
}

func (controller *UserController) Register(c echo.Context)error{
	ctx := c.Request().Context()

	var register request.Register

	err := c.Bind(&register)

	if err != nil{
		return controllers.NewErrorResponse(c, http.StatusInternalServerError,err)
	}
	user,err := controller.usecase.Register(*register.ToDomain(),ctx)
	if err != nil{
		return controllers.NewErrorResponse(c, http.StatusInternalServerError,err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomain(user))
}

func (controller *UserController) Login(c echo.Context)error{
	ctx := c.Request().Context()

	var login request.Login

	err := c.Bind(&login)

	if err != nil{
		return controllers.NewErrorResponse(c, http.StatusInternalServerError,err)
	}
	user,err := controller.usecase.Login(*login.ToDomain(),ctx)
	if err != nil{
		return controllers.NewErrorResponse(c, http.StatusInternalServerError,err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomain(user))
}

