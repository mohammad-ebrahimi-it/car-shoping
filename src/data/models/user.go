package models

type User struct {
	BaseModel
	Username  string `gorm:"type:string;size:20;not null;unique"`
	FirstName string `gorm:"type:string;size:15;null"`
	LastName  string `gorm:"type:string;size:25;null"`
	Mobile    string `gorm:"type:string;size:11;null;unique;default:null"`
	Email     string `gorm:"type:string;size:64;null;unique;default:null"`
	Password  string `gorm:"type:string;size:64; not null"`
	Enable    bool   `gorm:"default:true"`
	UserRoles *[]UserRole
}
