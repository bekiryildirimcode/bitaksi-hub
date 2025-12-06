package routes

import (
	"github.com/bekiryildirimcode/driver-service/handlers"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Driver *handlers.DriverHandler
}

func SetupRoutes(e *echo.Echo, h Handlers) {
	api := e.Group("/api/v1/driver")

	//Drivers
	api.GET("", h.Driver.GetAll)
	api.POST("", h.Driver.Create)
	api.GET("/nearby", h.Driver.GetNearby)
	api.PUT("/:id", h.Driver.Update)

}
