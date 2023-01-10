package dto

type AuthDTO struct {
	Token      string `json:"token"`
	RefreshKey string `json:"refresh_key"`
}
