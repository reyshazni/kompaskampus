package dto

type UniversityDTO struct {
	ID       int    `json:"id"`
	FullName string `json:"name"`
	UniCode  string `json:"code"`
}
