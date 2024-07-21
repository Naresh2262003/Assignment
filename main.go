package main

import (
	"parking_lot_service/db"
	"parking_lot_service/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db.Connect()
	// db.InitDB()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routes
	e.GET("/park/available", handler.GetAvailableSpots)
	e.POST("/park", handler.ParkVehicle)
	e.POST("/unpark", handler.UnparkVehicle)

	e.Logger.Fatal(e.Start(":8080"))
}
