package dto

type GetOtpRequest struct {
	MobileNumber string `json:"mobile_number" binding:"required,mobile,min=11,max=11"`
}

type TokenDetail struct {
	AccessToken            string `json:"accessToken"`
	RefreshToken           string `json:"refreshToken"`
	AccessTokenExpireTime  int64  `json:"AccessTokenExpireTime"`
	RefreshTokenExpireTime int64  `json:"RefreshTokenExpireTime"`
}

type RegisterUserByUsernameRequest struct {
	FirstName string `json:"firstName" binding:"required,min=3"`
	LastName  string `json:"lastName" binding:"required,min=3"`
	Username  string `json:"username" binding:"required,min=5"`
	Email     string `json:"email" binding:"required,email,min=6"`
	Password  string `json:"password" binding:"required,password,min=6"`
}

type LoginByMobileRequest struct {
	MobileNumber string `json:"mobileNumber" binding:"required,mobile,min=11,max=11"`
	Otp          string `json:"otp" binding:"required,min=6,max=6"`
}
type RegisterLoginByMobileNumber struct {
	MobileNumber string `json:"mobileNumber" binding:"required,mobile,min=11,max=11"`
	Otp          string `json:"otp" binding:"required,min=6,max=6"`
}

type LoginByUsernameRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}
