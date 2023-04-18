package utils

import "github.com/stanhoenson/krushr/internal/models"

func PointsOfInterestToDistance(pointsOfInterest *[]models.PointOfInterest) float64 {
	dereferencedPointsOfInterest := *pointsOfInterest
	distance := 0.0
	arrayLength := len(dereferencedPointsOfInterest)
	if arrayLength > 1 {
		distance = Haversine(dereferencedPointsOfInterest[0].Latitude, dereferencedPointsOfInterest[0].Longitude, dereferencedPointsOfInterest[arrayLength-1].Latitude, dereferencedPointsOfInterest[arrayLength-1].Longitude)
	}

	return distance
}
