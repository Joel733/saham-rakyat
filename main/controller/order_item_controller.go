package controller

import (
	"fmt"
	"saham-rakyat/main/request"
	"saham-rakyat/main/response"
	"saham-rakyat/main/service"

	"github.com/labstack/echo/v4"
)

type orderItemHandler struct {
	ordersItemService service.OrdersItemService
}

type OrderItemHandler interface{
	Create(c echo.Context) error
	FindAllOrderItems(c echo.Context) error
	Detail(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

func (h *orderItemHandler) Create(c echo.Context) error{
	var bodyRequest request.OrdersItemRequest

	if err := c.Bind(&bodyRequest); err != nil {
		return response.ErrBadRequest(c,err)
	}

	if(bodyRequest.Price==0){
		return response.ErrBadRequest(c,fmt.Errorf("price can't be empty "))
	}

	result, err := h.ordersItemService.Create(bodyRequest)
	if err != nil {
		return response.ErrInternalServer(c,err)
	}

	return response.Success(c, result)
}

func (h *orderItemHandler) FindAllOrderItems(c echo.Context) error{
	var bodyRequest request.OrdersItemRequest

	if err := c.Bind(&bodyRequest); err != nil {
		return response.ErrBadRequest(c,err)
	}

	resp, err := h.ordersItemService.FindAll(bodyRequest)
	if err != nil {
		return response.ErrInternalServer(c, err)
	}

	return response.Success(c, resp)
}

func (h *orderItemHandler) Detail(c echo.Context) error{
	
	var bodyRequest request.OrdersItemRequest

	if err := c.Bind(&bodyRequest); err != nil {
		return response.ErrBadRequest(c,err)
	}

	resp, err := h.ordersItemService.Detail(bodyRequest.ID)
	if err != nil {
		return response.ErrInternalServer(c, err)
	}

	return response.Success(c, resp)
}

func(h *orderItemHandler) Update(c echo.Context) error{

	var bodyRequest request.OrdersItemRequest

	if err := c.Bind(&bodyRequest); err != nil {
		return response.ErrBadRequest(c,err)
	}

	resp, err := h.ordersItemService.Update(bodyRequest)
	if err != nil {
		return response.ErrInternalServer(c, err)
	}

	return response.Success(c, resp)

}

func (h *orderItemHandler) Delete(c echo.Context) error{
	
	var bodyRequest request.OrdersItemRequest

	if err := c.Bind(&bodyRequest); err != nil {
		return response.ErrBadRequest(c,err)
	}

	resp, err := h.ordersItemService.Delete(bodyRequest)
	if err != nil {
		return response.ErrInternalServer(c, err)
	}

	return response.Success(c, resp)
}

