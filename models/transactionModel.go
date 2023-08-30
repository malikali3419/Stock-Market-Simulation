package models

import (
	"gorm.io/gorm"
)

type TransactionData struct {
	gorm.Model
	TransactionID     string
	Ticker            string
	TransactionType   string
	TransactionVolume int
	TransactionPrice  float32
	UserID            uint
	User              Users `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
