package models

type CarModelImage struct {
	BaseModel
	CarModel    CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	CarModelId  int
	Image       File `gorm:"foreignKey:ImageId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	ImageId     int
	IsMainImage bool
}
