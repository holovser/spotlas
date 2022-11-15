package utils

import (
	"github.com/paulsmith/gogeos/geos"
	"github.com/umahmood/haversine"
	"sort"
	"spotlas/src/model"
)

func ValidateGetSpotsParams(latitude string, longitude string,
	radius string, areaType string) bool {
	if latitude == "" || longitude == "" || radius == "" || areaType == "" {
		return false
	}
	return true
}

func SortByDistanceAndRank(spots []model.Spot, baseLong float64, baseLat float64) {
	bPointGeog := haversine.Coord{Lat: baseLat, Lon: baseLong}

	sort.Slice(spots, func(i, j int) bool {
		p1Geom, _ := geos.FromHex(spots[i].Coordinates)
		p1X, _ := p1Geom.X()
		p1Y, _ := p1Geom.Y()
		p1Geog := haversine.Coord{Lat: p1Y, Lon: p1X}

		_, d1 := haversine.Distance(bPointGeog, p1Geog)

		p2Geom, _ := geos.FromHex(spots[j].Coordinates)
		p2X, _ := p2Geom.X()
		p2Y, _ := p2Geom.Y()
		p2Geog := haversine.Coord{Lat: p2Y, Lon: p2X}

		_, d2 := haversine.Distance(bPointGeog, p2Geog)

		if d1-d2 < 0.05 {
			return spots[i].Rating < spots[j].Rating
		} else {
			return d1 < d2
		}
	})
}
