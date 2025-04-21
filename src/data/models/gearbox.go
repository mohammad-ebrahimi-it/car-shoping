package models

type Gearbox struct {
	BaseModel
	Name      string `gorm:"size=15;type=string;not null"`
	CarModels *[]CarModel
}
