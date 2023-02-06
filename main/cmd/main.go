package main

import (
	"log"
	"saham-rakyat/main/routing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	var err error

	err = godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routing.UserInitRoute(e)
	routing.OrderItemInitRoute(e)
	routing.OrderHistoriesInitRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}