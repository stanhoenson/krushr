package utils

import (
	"math"
)

func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371 // kilometers
	lat1Radians := lat1 * math.Pi / 180.0
	lat2Radians := lat2 * math.Pi / 180.0
	latDiffRadians := (lat2 - lat1) * math.Pi / 180.0
	lonDiffRadians := (lon2 - lon1) * math.Pi / 180.0

	a := math.Sin(latDiffRadians/2)*math.Sin(latDiffRadians/2) +
		math.Cos(lat1Radians)*math.Cos(lat2Radians)*
			math.Sin(lonDiffRadians/2)*math.Sin(lonDiffRadians/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}
