package models

type PropertyCategory struct {
	BaseModel
	Name       string      `gorm:"size:15;type:string;not null;unique"`
	Icon       string      `gorm:"size:1000;type:string;not null;unique"`
	Properties *[]Property `gorm:"foreignKey:CategoryId"`
}
