package routes

import (
	s "backend/server"
	"backend/server/handlers"
	"backend/services"
	"backend/stores"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ConfigureRoutes(server *s.Server) {
	stores := stores.New(server.DB)
	services := services.New(stores)
	handlers := handlers.New(services)

	server.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${path} (${remote_ip}) ${latency_human}\n",
		Output: server.Echo.Logger.Output(),
	}))
	server.Echo.Use(middleware.Recover())
	server.Echo.Pre(middleware.RemoveTrailingSlash())

	g := server.Echo.Group("api/v1")

	// Admin Endpoints
	// TODO: Configure auth middleware
	admin := g.Group("/admin")
	admin.GET("/user/:id", handlers.UserHanlder.GetUserById)
	admin.POST("/user", handlers.UserHanlder.CreateUser)
	admin.PUT("/user", handlers.UserHanlder.UpdateUserById)
	admin.DELETE("/user/:id", handlers.UserHanlder.DeleteUserById)

	// User Endpoints
	// TODO: Configure auth middleware
	user := g.Group("/user")
	user.GET("/:id", handlers.GetUserById)
	user.POST("/login", todo)
	user.POST("/register", todo)
	user.POST("/refresh", todo)

	restaurant := user.Group("/restaurant")
	restaurant.GET("/:id", todo)
	restaurant.GET("/orders", todo)
	restaurant.GET("/orders/:id", todo)
	restaurant.PUT("/orders", todo)
	restaurant.DELETE("/orders/:id", todo)

	// restaurant/categories CRUD
	restaurant.GET("/categories", todo)
	restaurant.POST("/categories", todo)
	restaurant.PUT("/categories", todo)
	restaurant.DELETE("/categories", todo)

	// restaurant/items CRUD
	restaurant.GET("/items", todo)
	restaurant.GET("/items/:id", todo)
	restaurant.POST("/items", todo)
	restaurant.PUT("/items", todo)
	restaurant.DELETE("/items", todo)

	// restaurant/tables CRUD
	restaurant.GET("/tables", todo)
	restaurant.GET("/tables/:id", todo)
	restaurant.POST("/tables", todo)
	restaurant.PUT("/tables", todo)
	restaurant.DELETE("/tables", todo)

	// Customer Endpoints
	// TODO: Configure auth middleware
	customer := g.Group("/customer")
	customer.GET("/", todo)
	customer.GET("/items", todo)
	customer.GET("/items/:id", todo)
	customer.GET("/orders", todo)
	customer.POST("/orders", todo)

}

func todo(c echo.Context) error {
	return c.JSON(http.StatusOK, "todo")
}
