package models

import "time"

type PersianYear struct {
	BaseModel
	PersianTitle  string    `gorm:"size:10;type:string;not null;unique"`
	Year          int       `gorm:"type:int;unique;index;not null"`
	StartAt       time.Time `gorm:"type:TIMESTAMP with time zone;not null;unique"`
	EndAt         time.Time `gorm:"type:TIMESTAMP with time zone;not null;unique"`
	CarModelYears *[]CarModelYear
}
