package routing

import (
	"saham-rakyat/main/config"
	"saham-rakyat/main/controller"
	"saham-rakyat/main/repository/mysql"
	"saham-rakyat/main/service"

	"github.com/labstack/echo/v4"
)
func UserInitRoute(e *echo.Echo){
	db, err := config.InitDatabase()
	if err != nil {
		panic(err)
	}

	userRepo := mysql.NewUserTable(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserHandler(userService)

	userRoute := e.Group("/user")
	userRoute.POST("/create", userController.Create)
	userRoute.GET("/list", userController.FindAllUsers)
	userRoute.GET("/detail", userController.Detail)
	userRoute.DELETE("/delete", userController.Delete)
	userRoute.PUT("/update", userController.Update)

	
}