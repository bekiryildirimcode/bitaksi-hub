package service

import (
	"math"
	"sort"

	"github.com/bekiryildirimcode/driver-service/model"
	"github.com/bekiryildirimcode/driver-service/util"
)

type driverService struct {
	repo model.DriverRepository
}

func NewDriverService(repo model.DriverRepository) model.DriverService {
	return &driverService{repo: repo}
}

func (s *driverService) CreateDriver(req *model.CreateDriverRequest) (string, error) {

	driver := &model.Driver{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Plate:     req.Plate,
		TaksiType: req.TaksiType,
		CarBrand:  req.CarBrand,
		CarModel:  req.CarModel,
		Location: model.Location{
			Lat: req.Lat,
			Lon: req.Lon,
		},
	}
	return s.repo.Create(driver)
}

func (s *driverService) UpdateDriver(id string, req *model.CreateDriverRequest) error {
	updates := map[string]interface{}{
		"firstName": req.FirstName,
		"lastName":  req.LastName,
		"location":  model.Location{Lat: req.Lat, Lon: req.Lon},
	}
	return s.repo.Update(id, updates)
}

func (s *driverService) ListDrivers(page, pageSize int) ([]model.Driver, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	return s.repo.List(page, pageSize)
}

func (s *driverService) FindNearby(req *model.NearbyDriverQuery) ([]model.NearbyDriverResponse, error) {
	drivers, err := s.repo.FindAllByType(req.TaksiType)
	if err != nil {
		return nil, err
	}

	var nearby []model.NearbyDriverResponse
	const radiusKm = 6.0

	for _, d := range drivers {
		dist := util.Haversine(req.Lat, req.Lon, d.Location.Lat, d.Location.Lon)

		if dist <= radiusKm {
			nearby = append(nearby, model.NearbyDriverResponse{
				FirstName:  d.FirstName,
				LastName:   d.LastName,
				Plate:      d.Plate,
				DistanceKm: math.Round(dist*100) / 100,
			})
		}
	}

	sort.Slice(nearby, func(i, j int) bool {
		return nearby[i].DistanceKm < nearby[j].DistanceKm
	})

	return nearby, nil
}
