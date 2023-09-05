package models

import (
	"gorm.io/gorm"
)

type TransactionType string

const (
	Buy  TransactionType = "buy"
	Sell TransactionType = "sell"
)

type TransactionData struct {
	gorm.Model
	TransactionID     string
	Ticker            string
	TransactionType   TransactionType
	TransactionVolume int
	TransactionPrice  float32
	UserID            uint
	User              Users `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
