package middlewares

import (
	"backend/stores"
	"backend/utils"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	store *stores.Stores
}

func NewAuthMw(store *stores.Stores) AuthMiddleware {
	return AuthMiddleware{
		store: store,
	}
}

func (m *AuthMiddleware) RestaurantAccess() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			restaurant_id := c.Param("restaurant_id")
			c.Set("restaurant_id", restaurant_id)

			token := c.Get("user").(*jwt.Token)
			claims := token.Claims.(*utils.JwtCustomClaims)

			user, err := m.store.User.GetById(claims.ID)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Access Denied")
			}

			// UserがこのRestaurantにアクセスできる権限があるかどうかを確認
			isOwner := false
			for _, r := range user.Restaurants {
				if restaurant_id == r.ID {
					isOwner = true
					break
				}
			}

			if !isOwner {
				return echo.NewHTTPError(http.StatusUnauthorized, "Access Denied")
			}
			// Set UserID in context
			c.Set("userID", user.ID)
			return next(c)
		}
	}
}
