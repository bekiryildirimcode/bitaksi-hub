package model

import (
	"time"
)

type Driver struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	FirstName string    `bson:"firstName" json:"firstName"`
	LastName  string    `bson:"lastName" json:"lastName"`
	Plate     string    `bson:"plate" json:"plate"`
	TaksiType string    `bson:"taksiType" json:"taksiType"`
	CarBrand  string    `bson:"carBrand" json:"carBrand"`
	CarModel  string    `bson:"carModel" json:"carModel"`
	Location  Location  `bson:"location" json:"location"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

type Location struct {
	Lat float64 `bson:"lat" json:"lat"`
	Lon float64 `bson:"lon" json:"lon"`
}

// DTO
type CreateDriverRequest struct {
	FirstName string  `json:"firstName" validate:"required"`
	LastName  string  `json:"lastName" validate:"required"`
	Plate     string  `json:"plate" validate:"required"`
	TaksiType string  `json:"taksiType" validate:"required"`
	CarBrand  string  `json:"carBrand"`
	CarModel  string  `json:"carModel"`
	Lat       float64 `json:"lat" validate:"required"`
	Lon       float64 `json:"lon" validate:"required"`
}

type NearbyDriverResponse struct {
	FirstName  string  `json:"firstName"`
	LastName   string  `json:"lastName"`
	Plate      string  `json:"plate"`
	DistanceKm float64 `json:"distanceKm"`
}

type NearbyDriverQuery struct {
	Lat       float64 `query:"lat" validate:"required"`
	Lon       float64 `query:"lon" validate:"required"`
	TaksiType string  `query:"taksiType" validate:"required"`
}

type DriverRepository interface {
	Create(driver *Driver) (string, error)
	Update(id string, updates map[string]interface{}) error
	List(page, pageSize int) ([]Driver, error)
	FindAllByType(taksiType string) ([]Driver, error)
}

type DriverService interface {
	CreateDriver(req *CreateDriverRequest) (string, error)
	UpdateDriver(id string, req *CreateDriverRequest) error
	ListDrivers(page, pageSize int) ([]Driver, error)
	FindNearby(req *NearbyDriverQuery) ([]NearbyDriverResponse, error)
}
