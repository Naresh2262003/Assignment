package models

import (
	"gorm.io/gorm"
)

type Tariff struct {
	gorm.Model
	ParkingLotID  uint    `json:"parking_lot_id"`
	VehicleTypeID uint    `json:"vehicle_type_id"`
	HourlyRate    float64 `json:"hourly_rate"`
	DailyRate     float64 `json:"daily_rate,omitempty"`
}
