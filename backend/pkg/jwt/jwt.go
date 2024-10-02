package jwt

import (
	config "auraskin/internal/configs/dev"
	"auraskin/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = user.ID
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["isAdmin"] = user.IsAdmin
	claims["exp"] = time.Now().Add(time.Hour * 168).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	cfg, err := config.Instance()
	if err != nil {
		return "", err
	}

	return token.SignedString([]byte(cfg.GetSecretKey()))
}