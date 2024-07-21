package handler

import (
	"math"
	"net/http"
	"time"

	"parking_lot_service/db"
	"parking_lot_service/models"

	"github.com/labstack/echo/v4"
)

func ParkVehicle(c echo.Context) error {
	type Request struct {
		VehicleNumber string `json:"vehicle_number"`
		VehicleTypeID uint   `json:"vehicle_type_id"`
		ParkingLotID  uint   `json:"parking_lot_id"`
	}

	var req Request
	err := c.Bind(&req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// checking for the availability of the spot
	var spot models.Spot
	db.DB.Where("parking_lot_id= ? AND vehicle_type_id = ? AND occupied_spots < total_spots", req.ParkingLotID, req.VehicleTypeID).First(&spot)
	if spot.ID == 0 {
		return c.JSON(http.StatusBadRequest, err)
	}

	// creating parking ticket
	ticket := models.ParkingTicket{
		VehicleNumber: req.VehicleNumber,
		VehicleTypeID: req.VehicleTypeID,
		ParkingLotID:  req.ParkingLotID,
		EntryTime:     time.Now(),
	}

	db.DB.Create(&ticket)

	// update Occupied Space
	db.DB.Model(&spot).Update("occupied_spots", spot.OccupiedSpots+1)

	return c.JSON(http.StatusOK, ticket)
}

func UnparkVehicle(c echo.Context) error {
	type Request struct {
		TicketID uint `json:"ticket_id"`
	}

	var req Request
	err := c.Bind(&req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var ticket models.ParkingTicket
	db.DB.First(&ticket, req.TicketID)
	if ticket.ID == 0 || ticket.ExitTime != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// calculating the fee
	exitTime := time.Now()
	duration := exitTime.Sub(ticket.EntryTime).Hours()
	roundedDuration := math.Ceil(duration) // Rounding up to the next whole hour

	var tariff models.Tariff
	db.DB.Where("parking_lot_id = ? AND vehicle_type_id= ?", ticket.ParkingLotID, ticket.VehicleTypeID).First(&tariff)

	var fee float64

	switch ticket.VehicleTypeID {
	case 7: // Motorcycles/scooters
		fee = roundedDuration * tariff.HourlyRate
	case 8: // Cars/SUVs
		if ticket.ParkingLotID == 5 { // Parking Lot A
			fee = roundedDuration * tariff.HourlyRate
		} else if ticket.ParkingLotID == 6 { // Parking Lot B
			if roundedDuration <= 1 {
				fee = 50
			} else {
				fee = tariff.HourlyRate + (roundedDuration-1)*tariff.HourlyRate
			}
		}
	case 9: // Buses/Trucks
		if ticket.ParkingLotID == 5 { // Parking Lot A
			if roundedDuration <= 24 {
				fee = roundedDuration * tariff.HourlyRate
			} else {
				days := math.Ceil(roundedDuration / 24)
				fee = days * tariff.DailyRate
			}
		} else if ticket.ParkingLotID == 6 { // Parking Lot B
			fee = roundedDuration * tariff.HourlyRate
		}
	default:
		return c.JSON(http.StatusBadRequest, "Invalid vehicle type")
	}

	// // update ticket with exit time and fee
	db.DB.Model(&ticket).Updates(map[string]interface{}{"exit_time": exitTime, "fee": fee})

	// Update occupied spots
	var spot models.Spot
	db.DB.Where("parking_lot_id = ? AND vehicle_type_id = ?", ticket.ParkingLotID, ticket.VehicleTypeID).First(&spot)
	db.DB.Model(&spot).Update("occupied_spots", spot.OccupiedSpots-1)

	return c.JSON(http.StatusOK, ticket)
}

func GetAvailableSpots(c echo.Context) error {
	type Request struct {
		ParkingLotID uint `json:"parking_lot_id"`
	}

	var req Request
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var spots []models.Spot
	db.DB.Where("parking_lot_id = ?", req.ParkingLotID).Find(&spots)

	return c.JSON(http.StatusOK, spots)
}
