package model

type Spot struct {
	Id          string  `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name        string  `json:"name"`
	Website     string  `json:"website"`
	Coordinates string  `json:"coordinates"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
}
