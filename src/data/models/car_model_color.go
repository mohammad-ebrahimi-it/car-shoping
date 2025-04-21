package models

type CarModelColor struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	CarModelId int
	Color      Color `gorm:"foreignKey:ColorId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	ColorId    int
}
