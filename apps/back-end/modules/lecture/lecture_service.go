package lecture

import (
	"FindMyDosen/model/dto"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AddLectureDTO struct {
	FullName     string `json:"fullName" validate:"required"`
	UniversityID int    `json:"university_id" validate:"required"`
	Faculty      string `json:"faculty" validate:"required"`
	Href         string `json:"href"`
}

func handleAddLecture(c echo.Context) error {
	lectureDTO := new(AddLectureDTO)
	if err := c.Bind(&lectureDTO); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BaseDTO{
			Message: "Cannot Bind JSON",
			Status:  http.StatusInternalServerError,
		})
	}

	validate := validator.New()
	if err := validate.Struct(lectureDTO); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			dto.BaseDTO{
				Status:  http.StatusBadRequest,
				Message: "Missing Field",
				Data:    err.Error(),
			},
		)
	}

	if err := addLecture(lectureDTO); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dto.BaseDTO{
				Status:  http.StatusInternalServerError,
				Message: "Cannot add lecture",
				Data:    err.Error(),
			},
		)
	}

	return c.JSON(http.StatusAccepted, dto.BaseDTO{Status: http.StatusAccepted, Message: "Lecture has been added into the database!"})
}
