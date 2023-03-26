package Authentication

import (
	"os"
	"time"

	"Academy/Models"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func GetToken(shop Models.Shop) (string, error) {

	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)

	claims["ShopID"] = shop.ID
	claims["exp"] = time.Now().Add(10 * time.Minute)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
