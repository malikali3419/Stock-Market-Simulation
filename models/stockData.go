package models

import "gorm.io/gorm"

type StockData struct {
	gorm.Model
	Ticker     string
	OpenPrice  float32
	ClosePrice float32
	High       float32
	Low        float32
	Volume     int
}
