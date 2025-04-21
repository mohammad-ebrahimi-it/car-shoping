package models

type Company struct {
	BaseModel
	Name      string  `gorm:"size=15;type=string;not null"`
	Country   Country `gorm:"foreignKey:CountryId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	CountryId int
	CarModels *[]CarModel
}
