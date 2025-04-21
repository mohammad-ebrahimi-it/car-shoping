package models

type CarModelYear struct {
	BaseModel
	CarModel             CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	CarModelId           int
	PersianYear          PersianYear `gorm:"foreignKey:PersianYearId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	PersianYearId        int
	CarModelPriceHistory *[]CarModelPriceHistory
}
