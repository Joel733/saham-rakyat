package controller

import (
	"saham-rakyat/main/request"
	"saham-rakyat/main/response"
	"saham-rakyat/main/service"

	"github.com/labstack/echo/v4"
)

type orderHistoriesHandler struct {
	orderHistoriesService service.OrdersHistoriesService
}

type OrderHistoriesHandler interface {
	Create(c echo.Context) error
	FindAllOrderHistories(c echo.Context) error
	Detail(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

func NewOrderHistoriesController(orderHistoriesService service.OrdersHistoriesService) OrderHistoriesHandler{
	return &orderHistoriesHandler{}
}

func (h *orderHistoriesHandler) Create(c echo.Context) error{
	var bodyRequest request.OrderHistoriesRequest

	if err := c.Bind(&bodyRequest); err != nil {
		return response.ErrBadRequest(c,err)
	}

	result, err := h.orderHistoriesService.Create(bodyRequest)
	if err != nil {
		return response.ErrInternalServer(c,err)
	}

	return response.Success(c, result)
}

func (h *orderHistoriesHandler) FindAllOrderHistories(c echo.Context) error{
	var bodyRequest request.FindAllOrderHistoriesItems

	if err := c.Bind(&bodyRequest); err != nil {
		return response.ErrBadRequest(c,err)
	}

	resp, err := h.orderHistoriesService.FindAll(bodyRequest)
	if err != nil {
		return response.ErrInternalServer(c, err)
	}

	return response.Success(c, resp)
}

func (h *orderHistoriesHandler) Detail(c echo.Context) error{
	
	var bodyRequest request.OrderHistoriesRequest

	if err := c.Bind(&bodyRequest); err != nil {
		return response.ErrBadRequest(c,err)
	}

	resp, err := h.orderHistoriesService.Detail(bodyRequest.ID)
	if err != nil {
		return response.ErrInternalServer(c, err)
	}

	return response.Success(c, resp)
}

func(h *orderHistoriesHandler) Update(c echo.Context) error{

	var bodyRequest request.OrderHistoriesRequest

	if err := c.Bind(&bodyRequest); err != nil {
		return response.ErrBadRequest(c,err)
	}

	resp, err := h.orderHistoriesService.Update(bodyRequest)
	if err != nil {
		return response.ErrInternalServer(c, err)
	}

	return response.Success(c, resp)

}

func (h *orderHistoriesHandler) Delete(c echo.Context) error{
	
	var bodyRequest request.OrderHistoriesRequest

	if err := c.Bind(&bodyRequest); err != nil {
		return response.ErrBadRequest(c,err)
	}

	resp, err := h.orderHistoriesService.Delete(bodyRequest)
	if err != nil {
		return response.ErrInternalServer(c, err)
	}

	return response.Success(c, resp)
}