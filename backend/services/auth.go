package services

import (
	"backend/config"
	"backend/stores"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type (
	AuthService interface {
		GenerateAccessToken(username string, password string, admin bool) (accessToken string, exp int64, err error)
	}

	authService struct {
		*config.Config
		*stores.Stores
	}
)

func (s *authService) GenerateAccessToken(username string, password string, admin bool) (accessToken string, exp int64, err error) {

	user, err := s.User.GetByUsername(username)
	if err != nil || (bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil) {
		return "", 0, err
	}

	expired := time.Now().Add(time.Hour * 72)

	claims := &config.JwtCustomClaims{
		ID:    user.ID,
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
