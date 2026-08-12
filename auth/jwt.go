package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/solace06/auth-service/internal/user"
)

func GenerateJWT(user user.User) (string, error) {

	secret := os.Getenv("JWT_SECRET")

	//claims
	claims := jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
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
