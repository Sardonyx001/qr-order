package config

import (
	"os"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/golang-jwt/jwt/v4"
)

type AuthConfig struct {
	AccessSecret  string
	RefreshSecret string
}

func LoadAuthConfig() AuthConfig {
	return AuthConfig{
		AccessSecret:  os.Getenv("ACCESS_SECRET"),
		RefreshSecret: os.Getenv("REFRESH_SECRET"),
	}
}

type BasicAuth struct {
	Username string `json:"username" validate:"required" example:"test_username"`
	Password string `json:"password" validate:"required" example:"test_password"`
}

func (ba BasicAuth) Validate() error {
	return validation.ValidateStruct(&ba,
		validation.Field(&ba.Username, validation.Length(8, 255)),
		validation.Field(&ba.Password, validation.Length(8, 255)),
	)
}

type JwtCustomClaims struct {
	ID    string `json:"id"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}
