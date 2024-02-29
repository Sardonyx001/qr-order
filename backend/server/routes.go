package server

import (
	"backend/middlewares"
	"backend/server/handlers"
	"backend/services"
	"backend/stores"
	"backend/utils"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ConfigureRoutes(server *Server) {
	stores := stores.New(server.DB)
	services := services.New(stores, server.Config)
	handlers := handlers.New(services)
	authMiddlware := middlewares.NewAuthMw(stores) // WORKSSSSS

	server.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${path} (${remote_ip}) ${latency_human}\n",
		Output: server.Echo.Logger.Output(),
	}))
	server.Echo.Use(middleware.Recover())
	server.Echo.Pre(middleware.RemoveTrailingSlash())

	g := server.Echo.Group("api/v1")
	jwtConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(utils.JwtCustomClaims)
		},
		SigningKey: []byte(server.Config.Auth.AccessSecret),
	}

	// Admin Endpoints
	admin := g.Group("/admin")
	admin.POST("/register", handlers.AdminHandler.CreateAdmin)
	admin.POST("/login", handlers.AuthHandler.LoginForAdmin)
	admin.Use(echojwt.WithConfig(jwtConfig))

	// User Endpoints
	users := g.Group("/users")
	users.POST("/register", handlers.UserHandler.CreateUser)
	users.POST("/login", handlers.AuthHandler.LoginForUser)

	users.Use(echojwt.WithConfig(jwtConfig))
	users.GET("/:id", handlers.UserHandler.GetUserById)
	users.PUT("/:id", handlers.UserHandler.UpdateUserById)
	users.DELETE("/:id", handlers.UserHandler.DeleteUserById)

	// Restaurant Endpoints
	restaurants := g.Group("/restaurants")
	restaurants.Use(echojwt.WithConfig(jwtConfig))
	restaurants.GET("", handlers.RestaurantHandler.GetRestaurants)
	restaurants.POST("", handlers.RestaurantHandler.CreateRestaurant)

	restaurants.Use(authMiddlware.RestaurantAccess())
	restaurants.GET("/:restaurant_id", handlers.RestaurantHandler.GetRestaurantById)
	restaurants.PUT("/:restaurant_id", handlers.RestaurantHandler.UpdateRestaurantName)

	// restaurant/:restaurant_id/orders CRUD
	orders := restaurants.Group("/:restaurant_id/orders")
	orders.GET("", todo)
	orders.POST("", todo)
	orders.GET("/:order_id", todo)
	orders.PUT("/:order_id", todo)
	orders.DELETE("/:order_id", todo)

	// restaurant/:restaurant_id/categories CRUD
	categories := restaurants.Group("/:restaurant_id/categories")
	categories.GET("", todo)
	categories.POST("", todo)
	categories.GET("/:category_id", todo)
	categories.PUT("/:category_id", todo)
	categories.DELETE("/:category_id", todo)

	// restaurant/:restaurant_id/items CRUD
	items := restaurants.Group("/:restaurant_id/items")
	items.GET("", handlers.ItemHandler.GetItems)
	items.POST("", handlers.ItemHandler.CreateItem)
	items.GET("/:item_id", handlers.ItemHandler.GetItemById)
	items.PUT("/:item_id", handlers.ItemHandler.UpdateItemById)
	items.DELETE("/:item_id", handlers.ItemHandler.DeleteItemById)

	// restaurant/:restaurant_id/tables CRUD
	tables := restaurants.Group("/:restaurant_id/tables")
	tables.GET("", todo)
	tables.POST("", todo)
	tables.GET("/:table_id", todo)
	tables.PUT("/:table_id", todo)
	tables.DELETE("/:table_id", todo)

	// Customer Endpoints
	customers := g.Group("/customers")
	customers.Use(echojwt.WithConfig(jwtConfig))

	customers.GET("", todo)
	customers.POST("", todo)
	customers.GET("/:customer_id/orders", todo)
	customers.POST("/:customer_id/orders", todo)
}

func todo(c echo.Context) error {
	return c.JSON(http.StatusOK, "Path: "+c.Path())
}
