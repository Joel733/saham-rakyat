package routing

import (
	"saham-rakyat/main/config"
	"saham-rakyat/main/controller"
	"saham-rakyat/main/repository/mysql"
	"saham-rakyat/main/service"

	"github.com/labstack/echo/v4"
)
func OrderItemInitRoute(e *echo.Echo){
	db, err := config.InitDatabase()
	if err != nil {
		panic(err)
	}

	orderItemRepo := mysql.NewOrderItemsTable(db)
	orderItemService := service.NewOrderItemsService(orderItemRepo)
	orderItemController := controller.NewOrderItemController(orderItemService)

	userRoute := e.Group("/orderItem")
	userRoute.POST("/create", orderItemController.Create)
	userRoute.GET("/list", orderItemController.FindAllOrderItems)
	userRoute.GET("/detail", orderItemController.Detail)
	userRoute.DELETE("/delete", orderItemController.Delete)
	userRoute.PUT("/update", orderItemController.Update)

	
}