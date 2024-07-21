package db

import (
	"fmt"
	"log"
	"parking_lot_service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := "host=db user=Naresh password=Naresh12345 dbname=parkingDB port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Define the initial data
	vehicleTypes := []models.VehicleType{
		{Type: "Motorcycle/Scooter", SpotSize: 1},
		{Type: "Car/SUV", SpotSize: 2},
		{Type: "Bus/Truck", SpotSize: 3},
	}

	parkingLots := []models.ParkingLot{
		{Name: "Parking Lot A"},
		{Name: "Parking Lot B"},
	}

	// Insert vehicle types
	for _, vehicleType := range vehicleTypes {
		if err := db.Create(&vehicleType).Error; err != nil {
			log.Fatal("Failed to insert vehicle types:", err)
		}
	}

	// Insert parking_lots
	for _, parkingLot := range parkingLots {
		if err := db.Create(&parkingLot).Error; err != nil {
			log.Fatal("Failed to insert parking lots:", err)
		}
	}

	// Get inserted parking lots and vehicle types
	var parkingLotA, parkingLotB models.ParkingLot
	var motorcycle, car, bus models.VehicleType

	db.Where("name = ?", "Parking Lot A").First(&parkingLotA)
	db.Where("name = ?", "Parking Lot B").First(&parkingLotB)
	db.Where("type = ?", "Motorcycle/Scooter").First(&motorcycle)
	db.Where("type = ?", "Car/SUV").First(&car)
	db.Where("type = ?", "Bus/Truck").First(&bus)

	// Insert spots
	spots := []models.Spot{
		{ParkingLotID: parkingLotA.ID, VehicleTypeID: motorcycle.ID, TotalSpots: 50},
		{ParkingLotID: parkingLotA.ID, VehicleTypeID: car.ID, TotalSpots: 30},
		{ParkingLotID: parkingLotA.ID, VehicleTypeID: bus.ID, TotalSpots: 20},
		{ParkingLotID: parkingLotB.ID, VehicleTypeID: motorcycle.ID, TotalSpots: 100},
		{ParkingLotID: parkingLotB.ID, VehicleTypeID: car.ID, TotalSpots: 80},
		{ParkingLotID: parkingLotB.ID, VehicleTypeID: bus.ID, TotalSpots: 40},
	}

	for _, spot := range spots {
		if err := db.Create(&spot).Error; err != nil {
			log.Fatal("Failed to insert spots:", err)
		}
	}

	// Insert tariffs
	tariffs := []models.Tariff{
		{ParkingLotID: parkingLotA.ID, VehicleTypeID: motorcycle.ID, HourlyRate: 5.00},
		{ParkingLotID: parkingLotA.ID, VehicleTypeID: car.ID, HourlyRate: 20.50},
		{ParkingLotID: parkingLotA.ID, VehicleTypeID: bus.ID, HourlyRate: 50.00, DailyRate: 500.00},
		{ParkingLotID: parkingLotB.ID, VehicleTypeID: motorcycle.ID, HourlyRate: 10.50},
		{ParkingLotID: parkingLotB.ID, VehicleTypeID: car.ID, HourlyRate: 25.00, DailyRate: 50.00},
		{ParkingLotID: parkingLotB.ID, VehicleTypeID: bus.ID, HourlyRate: 100.00},
	}

	for _, tariff := range tariffs {
		if err := db.Create(&tariff).Error; err != nil {
			log.Fatal("Failed to insert tariffs:", err)
		}
	}

	fmt.Println("Initial data inserted successfully!")
}
