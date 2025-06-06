package dto

type CreateUpdateCountryRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=20"`
}
type CreateCountryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CountryResponse struct {
	Id     int            `json:"id"`
	Name   string         `json:"name"`
	Cities []CityResponse `json:"cities"`
}
