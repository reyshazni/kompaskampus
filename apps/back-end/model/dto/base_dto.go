package dto

type BaseDTO struct {
	Status  int         `json:"status"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}
