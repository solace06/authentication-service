package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(id int64, role string) (string, error) {

	secret := os.Getenv("JWT_SECRET")

	//claims
	claims := jwt.MapClaims{
		"sub":  id,
		"role": role,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	//create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//sign token
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
