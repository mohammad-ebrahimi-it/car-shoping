package models

type Color struct {
	BaseModel
	Name           string `gorm:"size=15;type=string;not nul;unique"`
	HexCode        string `gorm:"size:7;type=string;not null;unique"`
	CarModelColors *[]CarModelColor
}
