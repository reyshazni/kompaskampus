package lecture

import (
	"FindMyDosen/model/dto"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type AddLectureDTO struct {
	FullName     string `json:"fullName" validate:"required"`
	UniversityID int    `json:"university_id" validate:"required"`
	Faculty      string `json:"faculty" validate:"required"`
	Href         string `json:"href"`
}

type LectureDTO struct {
	Id         uint   `json:"id"`
	FullName   string `json:"full_name"`
	University string `json:"university"`
	Faculty    string `json:"faculty"`
	Href       string `json:"href"`
}

type GetLectureQuery struct {
	Search string `query:"search"`
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
}

func handleGetLecture(c echo.Context) error {
	query := new(GetLectureQuery)
	if err := c.Bind(query); err != nil {
		return c.JSON(
			http.StatusNotFound,
			dto.BaseDTO{
				Status:  http.StatusNotFound,
				Message: "Lecture not found",
			})
	}
	result, err := getLecture(query)
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			dto.BaseDTO{
				Status:  http.StatusNotFound,
				Message: "Lecture not found",
			})
	}
	return c.JSON(
		http.StatusOK, dto.BaseDTO{
			Status:  http.StatusOK,
			Message: "Found!",
			Data:    result,
		})
}

func handleGetLectureByID(c echo.Context) error {
	id := c.Param("id")
	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			dto.BaseDTO{
				Status:  http.StatusNotFound,
				Message: "Lecture not found",
			})
	}
	dtoResult, err := getLectureById(uint(i))
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			dto.BaseDTO{
				Status:  http.StatusNotFound,
				Message: "Lecture not found",
			})
	}
	return c.JSON(
		http.StatusOK, dto.BaseDTO{
			Status:  http.StatusOK,
			Message: "Found!",
			Data:    dtoResult,
		})
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
