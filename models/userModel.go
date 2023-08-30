package models

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	UserID   string
	Username string
	Password string
	Balance  float32
	//Transactions []TransactionData
}
type JWTClaims struct {
	jwt.StandardClaims
	UserID uint `json:"user_id"`
}
