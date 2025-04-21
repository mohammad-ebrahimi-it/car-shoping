package models

import "time"

type CarModelPriceHistory struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	CarModelId int
	Price      float64   `gorm:"type:decimal(10,2);not null"`
	PriceAt    time.Time `gorm:"type:TIMESTAMP with time zone;not null"`
}
