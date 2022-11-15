package services

import (
	"gorm.io/gorm"
	"spotlas/src/model"
	"spotlas/src/utils"
)

type SpotsService struct {
	DB *gorm.DB
}

func (service SpotsService) GetSpotsInCircle(longitude float64, latitude float64,
	radius float64) []model.Spot {
	spots := make([]model.Spot, 0)
	service.DB.Raw(
		"select * from get_spots_in_circle(?, ?, ?)", longitude, latitude, radius,
	).Scan(&spots)

	utils.SortByDistanceAndRank(spots, longitude, latitude)
	return spots
}

func (service SpotsService) GetSpotsInSquare(longitude float64, latitude float64,
	radius float64) []model.Spot {
	spots := make([]model.Spot, 0)
	service.DB.Raw(
		"select * from get_spots_in_square(?, ?, ?)", longitude, latitude, radius,
	).Scan(&spots)

	utils.SortByDistanceAndRank(spots, longitude, latitude)
	return spots
}
