package models

import (
	"time"

	"gorm.io/gorm"
)

type ParkingTicket struct {
	gorm.Model
	VehicleNumber string     `json:"vehicle_number"`
	VehicleTypeID uint       `json:"vehicle_type_id"`
	ParkingLotID  uint       `json:"parking_lot_id"`
	EntryTime     time.Time  `json:"entry_time"`
	ExitTime      *time.Time `json:"exit_time"`
	Fee           float64    `json:"fee"`
}
