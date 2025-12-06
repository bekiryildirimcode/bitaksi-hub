package main

import (
	"net/http"

	"github.com/bekiryildirimcode/driver-gateway/handlers"
	"github.com/bekiryildirimcode/driver-gateway/middlewares"
	"github.com/bekiryildirimcode/driver-gateway/proxy"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middlewares.RequestLogger)
	limiter := middleware.RateLimiter(
		middleware.NewRateLimiterMemoryStore(5), // saniyede 5 istek/IP
	)
	//swager
	e.File("/docs/openapi.yaml", "docs/openapi.yaml")
	e.File("/docs", "docs/swagger.html")

	e.File("/swagger", "docs/swagger.html")
	e.POST("/login", handlers.Login, limiter)
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/swagger")
	})

	r := e.Group("/api/v1/")
	r.Use(middlewares.JWTMiddleware)
	r.Use(middlewares.RequestLogger)
	r.Use(limiter)
	driver := proxy.NewReverseProxy("http://driver-service:8000")
	r.Any("*", driver)

	e.Logger.Fatal(e.Start(":8080"))
}
