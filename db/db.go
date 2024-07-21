package db

import (
	"log"

	"parking_lot_service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=db user=Naresh password=Naresh12345 dbname=parkingDB sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(&models.ParkingLot{}, &models.VehicleType{}, &models.Spot{}, &models.ParkingTicket{}, &models.Tariff{})

	DB = database
}
