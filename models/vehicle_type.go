package models

import "gorm.io/gorm"

type VehicleType struct {
	gorm.Model
	Type     string `json:"type"`
	SpotSize int    `json:"spot_size"`
}
