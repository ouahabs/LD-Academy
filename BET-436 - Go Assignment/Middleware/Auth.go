package Middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func VerifyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")

		if token == "" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Please login"})
			return
		}

		newToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("WRONG METHOD: %v", t.Header["alg"])
			}

			return secretKey, nil
		})

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error, please try again"})
			return
		}

		if claims, ok := newToken.Claims.(jwt.MapClaims); ok && newToken.Valid {
			c.Set("ShopID", claims["ShopID"])
			c.Next()
		} else {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Please login."})
		}
	}
}
