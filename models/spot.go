package models

import "gorm.io/gorm"

type Spot struct {
	gorm.Model
	ParkingLotID  uint `json:"parking_lot_id"`
	VehicleTypeID uint `json:"vehicle_type_id"`
	TotalSpots    int  `json:"total_spots"`
	OccupiedSpots int  `json:"occupied_spots"`
}
