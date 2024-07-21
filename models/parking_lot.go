package models

import "gorm.io/gorm"

type ParkingLot struct {
	gorm.Model
	Name  string `json:"name"`
	Spots []Spot `json:"spots"`
}
