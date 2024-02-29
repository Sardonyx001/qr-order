package services

import (
	"backend/config"
	"backend/stores"
	"backend/utils"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type (
	AuthService interface {
		GenerateAccessToken(id string, admin bool) (accessToken string, exp int64, err error)
	}

	authService struct {
		*config.Config
		*stores.Stores
	}
)

func (s *authService) GenerateAccessToken(id string, admin bool) (accessToken string, exp int64, err error) {
	expired := time.Now().Add(time.Hour * 72)

	claims := &utils.JwtCustomClaims{
		ID:    id,
		Admin: admin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expired),
		},
	}

	exp = expired.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err = token.SignedString([]byte(s.Auth.AccessSecret))
	if err != nil {
		return "", 0, err
	}

	return accessToken, exp, err
}
