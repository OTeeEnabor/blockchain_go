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
	ContractId *ContractID
	ContractIdString string
	GasUsed uint64
	TransactionId string
	ChargeFee string
	PayerAccount string
	Status string
	ReceiptPDF string
	CreatedAt time.Time
}