package controller

import (
	"saham-rakyat/main/request"
	"saham-rakyat/main/response"
	"saham-rakyat/main/service"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService service.UserService
}

type UserHandler interface{
	Create(c echo.Context) error
	FindAllUsers(c echo.Context) error
	Detail(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

func NewUserHandler(userService service.UserService) UserHandler{
	return &userHandler{userService}
}

func (h *userHandler) Create(c echo.Context) error{
	var bodyRequest request.UsersRequest

	if err := c.Bind(&bodyRequest); err != nil {
		return response.ErrBadRequest(c,err)
	}

	result, err := h.userService.Create(bodyRequest)
	if err != nil {
		return response.ErrInternalServer(c,err)
	}

	return response.Success(c, result)
}

func (h *userHandler) FindAllUsers(c echo.Context) error{
	var bodyRequest request.FindAllUsersRequest

	if err := c.Bind(&bodyRequest); err != nil {
		return response.ErrBadRequest(c,err)
	}

	resp, err := h.userService.FindAll(bodyRequest)
	if err != nil {
		return response.ErrInternalServer(c, err)
	}

	return response.Success(c, resp)
}

func (h *userHandler) Detail(c echo.Context) error{
	
	var bodyRequest request.UsersRequest

	if err := c.Bind(&bodyRequest); err != nil {
		return response.ErrBadRequest(c,err)
	}

	resp, err := h.userService.Detail(bodyRequest.ID)
	if err != nil {
		return response.ErrInternalServer(c, err)
	}

	return response.Success(c, resp)
}

func(h *userHandler) Update(c echo.Context) error{

	var bodyRequest request.UsersRequest

	if err := c.Bind(&bodyRequest); err != nil {
		return response.ErrBadRequest(c,err)
	}

	resp, err := h.userService.Update(bodyRequest)
	if err != nil {
		return response.ErrInternalServer(c, err)
	}

	return response.Success(c, resp)

}

func (h *userHandler) Delete(c echo.Context) error{
	
	var bodyRequest request.UsersRequest

	if err := c.Bind(&bodyRequest); err != nil {
		return response.ErrBadRequest(c,err)
	}

	resp, err := h.userService.Delete(bodyRequest)
	if err != nil {
		return response.ErrInternalServer(c, err)
	}

	return response.Success(c, resp)
}