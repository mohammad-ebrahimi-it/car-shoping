package models

type CarModelProperty struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	CarModelId int
	Property   Property `gorm:"foreignKey:PropertyId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	PropertyId int
	Value      string `gorm:"size:100;type:string;not null"`
}
