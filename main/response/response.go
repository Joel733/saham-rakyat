package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Success(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

func ErrBadRequest(c echo.Context, err error) error {
	errResp := buildError(err)
	return c.JSON(http.StatusBadRequest, errResp)
}

func ErrInternalServer(c echo.Context, err error) error {
	errResp := buildError(err )
	return c.JSON(http.StatusInternalServerError, errResp)
}

func ErrNotFound(c echo.Context, err error) error {
	errResp := buildError(err )
	return c.JSON(http.StatusNotFound, errResp)
}

func buildError(err error) *ErrorResponse {
	return &ErrorResponse{
		Message: err.Error(),
	}
}
