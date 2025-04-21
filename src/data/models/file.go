package models

type File struct {
	BaseModel
	Name        string `gorm:"size=100;type=string;not nul;"`
	Directory   string `gorm:"size=100;type=string;not nul;"`
	Description string `gorm:"size=500;type=string;not nul;"`
	MimeType    string `gorm:"size=20;type=string;not nul;"`
}
