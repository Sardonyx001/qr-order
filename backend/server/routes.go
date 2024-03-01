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
	h := handlers.New(services)
	authMiddlware := middlewares.NewAuthMw(stores)

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
	admin.POST("/register", h.CreateAdmin)
	admin.POST("/login", h.LoginForAdmin)
	admin.Use(echojwt.WithConfig(jwtConfig))

	// User Endpoints
	users := g.Group("/users")
	users.POST("/register", h.CreateUser)
	users.POST("/login", h.LoginForUser)

	users.Use(echojwt.WithConfig(jwtConfig))
	users.GET("/:id", h.GetUserById)
	users.PUT("/:id", h.UpdateUserById)
	users.DELETE("/:id", h.DeleteUserById)

	// Restaurant Endpoints
	restaurants := g.Group("/restaurants")
	restaurants.Use(echojwt.WithConfig(jwtConfig))
	restaurants.GET("", h.GetRestaurants)
	restaurants.POST("", h.CreateRestaurant)

	restaurants.Use(authMiddlware.RestaurantAccess())
	restaurants.GET("/:restaurant_id", h.GetRestaurantById)
	restaurants.PUT("/:restaurant_id", h.UpdateRestaurant)

	// restaurant/:restaurant_id/categories CRUD
	categories := restaurants.Group("/:restaurant_id/categories")
	categories.GET("", h.GetCategories)
	categories.POST("", h.CreateCategory)
	categories.GET("/:category_id", h.GetCategoryById)
	categories.PUT("/:category_id", h.UpdateCategoryById)
	categories.DELETE("/:category_id", h.DeleteCategoryById)

	// restaurant/:restaurant_id/items CRUD
	items := restaurants.Group("/:restaurant_id/items")
	items.GET("", h.GetItems)
	items.POST("", h.CreateItem)
	items.GET("/:item_id", h.GetItemById)
	items.PUT("/:item_id", h.UpdateItemById)
	items.DELETE("/:item_id", h.DeleteItemById)

	// restaurant/:restaurant_id/tables CRUD
	tables := restaurants.Group("/:restaurant_id/tables")
	tables.GET("", h.GetTables)
	tables.POST("", h.CreateTable)
	tables.GET("/:table_id", h.GetTableById)
	tables.PUT("/:table_id", h.UpdateTableById)
	tables.DELETE("/:table_id", h.DeleteItemById)

	// restaurant/:restaurant_id/orders CRUD
	orders := restaurants.Group("/:restaurant_id/orders")
	orders.GET("", h.GetOrders)
	orders.POST("", h.CreateOrder)
	orders.GET("/:order_id", h.GetOrderById)
	orders.DELETE("/:order_id", h.DeleteOrderById)

	orderItems := restaurants.Group("/:restaurant_id/orderItems")
	orderItems.GET("/:order_item_id", h.GetOrderItemById)
	orderItems.PUT("/:order_item_id", h.UpdateOrderItemById)

	// Customer Endpoints
	customers := g.Group("/customers")
	customers.Use(echojwt.WithConfig(jwtConfig))

	customers.GET("", h.GetCustomerById)
	customers.POST("", h.CreateCustomer)
	customers.GET("/:customer_id/orders", h.GetOrdersForCustomer)
	customers.POST("/:customer_id/orders", todo)
}

func todo(c echo.Context) error {
	return c.JSON(http.StatusOK, "Path: "+c.Path())
}
