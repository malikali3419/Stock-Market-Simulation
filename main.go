package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"stock_market_simulation/m/controllers"
	"stock_market_simulation/m/initializers"
	"stock_market_simulation/m/models"
	"strings"
)

func init() {
	initializers.LoadEnviromentalVariables()
	initializers.ConnectToDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/users", controllers.ResgisterUser)
	r.POST("/login", controllers.LoginUser)
	r.Use(AuthMiddleware())
	r.GET("/allUsers", controllers.GetAllUsers)
	r.GET("user/:username", controllers.GetUser)
	r.POST("/stocks", controllers.AddStocks)
	r.GET("/stocks", controllers.GetAllStocks)
	r.GET("/stock/:ticker", controllers.GetOneStock)
	r.POST("/transactions", controllers.DoTransaction)
	r.GET("/transactions/:user_id", controllers.GetAllTransactionOFUser)
	r.GET("/transactions/:user_id/:start_time/:end_time", controllers.GetTransactionDataBetweenTime)

	err := r.Run()
	if err != nil {
		return
	}
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		// Split the Authorization header to extract the token
		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Malformed token"})
			c.Abort()
			return
		}
		tokenString := bearerToken[1] // Extract the actual token from the "Bearer <token>" format
		token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Make sure that the token method conforms to "SigningMethodHMAC"
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})

		if _, ok := token.Claims.(*models.JWTClaims); !ok || !token.Valid {
			var validationError *jwt.ValidationError
			if errors.As(err, &validationError) {
				if validationError.Errors&jwt.ValidationErrorMalformed != 0 {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Malformed token"})
				} else if validationError.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is either expired or not active yet"})
				} else {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
				}
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
			}
			c.Abort()
			return
		}

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		claims := token.Claims.(*models.JWTClaims)
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}
