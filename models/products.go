package models

import (
	"time"

	"gorm.io/gorm"
)
// database model
type Products struct {
	gorm.Model
	ProductID string
	Colour    string
	Quantity  int64
	Timestamp time.Time
}