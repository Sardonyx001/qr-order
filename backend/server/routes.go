package server

import (
	"backend/config"
	"backend/server/handlers"
	"backend/services"
	"backend/stores"
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

	server.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${path} (${remote_ip}) ${latency_human}\n",
		Output: server.Echo.Logger.Output(),
	}))
	server.Echo.Use(middleware.Recover())
	server.Echo.Pre(middleware.RemoveTrailingSlash())

	g := server.Echo.Group("api/v1")
	jwtConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(config.JwtCustomClaims)
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

	restaurants := g.Group("/restaurants")
	restaurants.Use(echojwt.WithConfig(jwtConfig))

	restaurants.GET("", handlers.RestaurantHandler.GetRestaurants)
	restaurants.GET("/:restaurant_id", todo)
	restaurants.POST("", handlers.RestaurantHandler.CreateRestaurant)

	restaurants.GET("/orders", todo)
	restaurants.GET("/orders/:order_id", todo)
	restaurants.POST("/orders/:order_id", todo)
	restaurants.PUT("/orders/:order_id", todo)
	restaurants.DELETE("/orders/:order_id", todo)

	// restaurant/:restaurant_id/categories CRUD
	restaurants.GET("/categories", todo)
	restaurants.GET("/categories/:category_id", todo)
	restaurants.POST("/categories/:category_id", todo)
	restaurants.PUT("/categories/:category_id", todo)
	restaurants.DELETE("/categories/:category_id", todo)

	// restaurant/:restaurant_id/items CRUD
	restaurants.GET("/items", todo)
	restaurants.POST("/items", todo)
	restaurants.GET("/items/:item_id", todo)
	restaurants.PUT("/items/:item_id", todo)
	restaurants.DELETE("/items/:item_id", todo)

	// restaurant/:restaurant_id/tables CRUD
	restaurants.GET("/tables", todo)
	restaurants.POST("/tables", todo)
	restaurants.GET("/tables/:table_id", todo)
	restaurants.PUT("/tables/:table_id", todo)
	restaurants.DELETE("/tables/:table_id", todo)

	// Customer Endpoints
	customers := g.Group("/customers")
	customers.Use(echojwt.WithConfig(jwtConfig))

	customers.GET("", todo)
	customers.GET("/:customer_id/orders", todo)
	customers.POST("/:customer_id/orders", todo)
}

func todo(c echo.Context) error {
	return c.JSON(http.StatusOK, "Path: "+c.Path())
}
