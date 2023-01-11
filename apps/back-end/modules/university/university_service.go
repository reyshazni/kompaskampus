package university

import (
	"FindMyDosen/model/dto"
	"FindMyDosen/model/query_param"
	"github.com/labstack/echo/v4"
	"net/http"
)

func handleGetUniversityData(c echo.Context) error {
	query := new(query_param.UniversityQuery)
	if err := c.Bind(query); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.BaseDTO{
			Status:  http.StatusInternalServerError,
			Message: "Bad Request",
			Data:    err.Error(),
		})
	}
	if query.Page == 0 {
		query.Page = 1
	}
	if query.Limit == 0 {
		query.Limit = 10
	}

	if query.Name != nil {
		result, err := getUniversityByName(query)
		if err != nil || len(result) == 0 || result == nil {
			print("error!")
			return c.JSON(
				http.StatusNotFound, dto.BaseDTO{
					Status:  http.StatusNotFound,
					Message: "University Not Found!",
				})
		}
		return c.JSON(
			http.StatusOK, dto.BaseDTO{
				Status:  http.StatusOK,
				Message: "Ok",
				Data:    result,
			})
	}

	result, err := getUniversity(query.Page, query.Limit)
	if err != nil || len(result) == 0 || result == nil {
		print("error!")
		return c.JSON(
			http.StatusNotFound, dto.BaseDTO{
				Status:  http.StatusNotFound,
				Message: "University Not Found!",
			})
	}
	return c.JSON(
		http.StatusOK, dto.BaseDTO{
			Status:  http.StatusOK,
			Message: "Ok",
			Data:    result,
		})
}
