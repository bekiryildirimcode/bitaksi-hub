package util

import (
	"math"
)

const R = 6371.0

func toRadians(degree float64) float64 {
	return degree * math.Pi / 180
}

// The Haversine formula gives the distance between two points on a sphere using their longitude and latitude.
func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	rLat1 := toRadians(lat1)
	rLon1 := toRadians(lon1)
	rLat2 := toRadians(lat2)
	rLon2 := toRadians(lon2)

	dLat := rLat2 - rLat1
	dLon := rLon2 - rLon1

	// Haversine formul: a = sin²(Δφ/2) + cos(φ₁) ⋅ cos(φ₂) ⋅ sin²(Δλ/2)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(rLat1)*math.Cos(rLat2)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c

	return distance
}
