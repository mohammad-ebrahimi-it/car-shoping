package models

type CarModel struct {
	BaseModel
	Name               string  `gorm:"size=15;type=string;not null"`
	Company            Company `gorm:"foreignKey:CompanyId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	CompanyId          int
	CarType            CarType `gorm:"foreignKey:CarTypeId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	CarTypeId          int
	Gearbox            Gearbox `gorm:"foreignKey:GearboxId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	GearboxId          int
	CarModelColors     *[]CarModelColor
	CarModelYears      *[]CarModelYear
	CarModelProperties *[]CarModelProperty
	CarModelImages     *[]CarModelImage
	CarModelComments   *[]CarModelComment
}
