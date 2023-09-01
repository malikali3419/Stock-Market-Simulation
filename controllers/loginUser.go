package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"stock_market_simulation/m/initializers"
	"stock_market_simulation/m/models"
	"time"
)

func init() {
	initializers.LoadEnviromentalVariables()
	initializers.ConnectToDatabase()
}

func LoginUser(c *gin.Context) {
	var user struct {
		Username string
		Password string
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var found_user models.Users
	initializers.DB.Where("username = ? ", user.Username).First(&found_user)

	if err := bcrypt.CompareHashAndPassword([]byte(found_user.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	claims := models.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
		UserID: found_user.ID,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
