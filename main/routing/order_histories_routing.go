package routing

import (
	"saham-rakyat/main/config"
	"saham-rakyat/main/controller"
	"saham-rakyat/main/repository/mysql"
	"saham-rakyat/main/service"

	"github.com/labstack/echo/v4"
)
func OrderHistoriesInitRoute(e *echo.Echo){
	db, err := config.InitDatabase()
	if err != nil {
		panic(err)
	}

	orderHistoriesRepo := mysql.NewOrderHistoriesTable(db)
	orderHistoriesService := service.NewOrderHistoriesService(orderHistoriesRepo)
	orderHistoriesController := controller.NewOrderHistoriesController(orderHistoriesService)

	userRoute := e.Group("/orderHistories")
	userRoute.POST("/create", orderHistoriesController.Create)
	userRoute.GET("/list", orderHistoriesController.FindAllOrderHistories)
	userRoute.GET("/detail", orderHistoriesController.Detail)
	userRoute.DELETE("/delete", orderHistoriesController.Delete)
	userRoute.PUT("/update", orderHistoriesController.Update)

	
}