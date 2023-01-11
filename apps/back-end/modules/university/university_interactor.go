package university

import (
	"FindMyDosen/database"
	"FindMyDosen/model/dto"
	"FindMyDosen/model/entity"
	"FindMyDosen/model/query_param"
)

func getUniversity(page int, limit int) ([]dto.UniversityDTO, error) {
	db := database.Paginate(page, limit)
	var universities []entity.University
	err := db.Order("id asc").Find(&universities).Error
	if err != nil {
		return nil, err
	}
	result := mapEntity(universities)
	return result, nil
}

func mapEntity(universities []entity.University) []dto.UniversityDTO {
	var result []dto.UniversityDTO
	for _, university := range universities {
		result = append(result, dto.UniversityDTO{
			ID:       university.ID,
			FullName: university.FullName,
			UniCode:  university.UniCode,
		})
	}
	return result
}

func getUniversityByName(query *query_param.UniversityQuery) ([]dto.UniversityDTO, error) {
	var universities []entity.University
	println("BY NAME")
	db := database.Paginate(query.Page, query.Limit)
	err := db.Where("full_name LIKE ?", "%"+*query.Name+"%").Order("id asc").Find(&universities).Error
	if err != nil {
		return nil, err
	}
	result := mapEntity(universities)
	return result, nil
}
