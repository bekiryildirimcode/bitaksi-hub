package main

import (
	"log"

	"github.com/bekiryildirimcode/driver-service/database"
	"github.com/bekiryildirimcode/driver-service/handlers"
	"github.com/bekiryildirimcode/driver-service/repository"
	"github.com/bekiryildirimcode/driver-service/routes"
	"github.com/bekiryildirimcode/driver-service/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err := database.ConnectMongoDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	repo := repository.NewMongoRepository(db)
	driverService := service.NewDriverService(repo)
	driverHandler := handlers.NewDriverHandler(driverService)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	routes.SetupRoutes(e, routes.Handlers{
		Driver: driverHandler,
	})

	e.Logger.Fatal(e.Start(":8000"))
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
