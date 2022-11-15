package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"net/http"
	"spotlas/src/http"
	"spotlas/src/model"
	"spotlas/src/services"
	"spotlas/src/utils"
	"strconv"
)

type Handler struct {
	SService services.SpotsService
}

func main() {
	router := gin.Default()
	handler := Handler{}
	handler.SService.DB = getDbConnection()

	router.GET("/spots", func(context *gin.Context) {
		latitudeStr := context.Query("latitude")
		longitudeStr := context.Query("longitude")
		radiusStr := context.Query("radius")
		areaTypeStr := context.Query("type")

		latitude, _ := strconv.ParseFloat(latitudeStr, 64)
		longitude, _ := strconv.ParseFloat(longitudeStr, 64)
		radius, _ := strconv.ParseFloat(radiusStr, 64)

		areaType := request.AreaType(areaTypeStr)
		//json.Unmarshal([]byte(context.Query("type")), &areaType)
		if !utils.ValidateGetSpotsParams(latitudeStr, longitudeStr, radiusStr, areaTypeStr) {
			context.JSON(400, gin.H{"code": "BAD REQUEST"})
			return
		} else {
			spots := make([]model.Spot, 0)
			switch areaType {
			case request.Circle:
				spots = handler.SService.GetSpotsInCircle(longitude, latitude, radius)
				context.IndentedJSON(http.StatusOK, spots)
				return
			case request.Square:
				spots = handler.SService.GetSpotsInSquare(longitude, latitude, radius)
				context.IndentedJSON(http.StatusOK, spots)
				return
			}
			context.JSON(400, gin.H{"code": "BAD REQUEST"})
		}
	})
	router.Run("localhost:8080")
}

func getDbConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}
	return db
}
