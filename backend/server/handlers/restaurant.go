package handlers

import (
	"backend/config"
	"backend/logger"
	"backend/models"
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type (
	RestaurantHandler interface {
		GetRestaurants(c echo.Context) error
		GetRestaurantById(c echo.Context) error
		CreateRestaurant(c echo.Context) error
		UpdateRestaurantName(c echo.Context) error
		DeleteRestaurant(c echo.Context) error
	}

	restaurantHandler struct {
		r services.RestaurantService
		u services.UserService
		c services.CategoryService
	}
)

func (h *restaurantHandler) GetRestaurants(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*config.JwtCustomClaims)
	user, err := h.u.GetUserById(claims.ID)
	if err != nil {
		logger.Error(err.Error())
		return c.JSON(http.StatusForbidden, "Invalid credentials")
	}

	return c.JSON(http.StatusOK, user.Restaurants)
}

func (h *restaurantHandler) GetRestaurantById(c echo.Context) error {
	restaurant_id := c.Param("restaurant_id")

	restaurant, err := h.r.GetRestaurantById(restaurant_id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, restaurant)
}

func (h *restaurantHandler) CreateRestaurant(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*config.JwtCustomClaims)
	user, err := h.u.GetUserById(claims.ID)
	if err != nil {
		return c.JSON(http.StatusForbidden, "Invalid credentials")
	}

	restaurant := &models.Restaurant{}
	if err := c.Bind(restaurant); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request form")
	}

	restaurant.Users = append(restaurant.Users, user)
	user.Restaurants = append(user.Restaurants, restaurant)

	createdRestaurant, err := h.r.CreateRestaurant(restaurant)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusCreated, createdRestaurant)
}

func (h *restaurantHandler) UpdateRestaurantName(c echo.Context) error {
	restaurant_id := c.Param("restaurant_id")

	modifiedRestaurant := &models.Restaurant{}
	if err := c.Bind(modifiedRestaurant); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request form")
	}

	var restaurant *models.Restaurant
	restaurant, err := h.r.GetRestaurantById(restaurant_id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Restaurant not found")
	}
	restaurant.Name = modifiedRestaurant.Name

	updatedRestaurant, err := h.r.UpdateRestaurant(restaurant)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusCreated, updatedRestaurant)
}

func (h *restaurantHandler) DeleteRestaurant(c echo.Context) (err error) {
	restaurant_id := c.Param("id")

	err = h.r.DeleteRestaurant(restaurant_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Error{Message: "ID is Invalid"})
	}
	return err
}
