package university

import (
	"FindMyDosen/database"
	"FindMyDosen/model/dto"
	"FindMyDosen/model/entity"
	"FindMyDosen/model/query_param"
)

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

func getUniversity(query *query_param.UniversityQuery) ([]dto.UniversityDTO, error) {
	var universities []entity.University
	println("BY NAME")
	db := database.Paginate(query.Page, query.Limit)
	err := db.Where("full_name LIKE ? OR uni_code LIKE ?", "%"+query.Search+"%", "%"+query.Search+"%").Order("id asc").Find(&universities).Error
	if err != nil {
		return nil, err
	}
	result := mapEntity(universities)
	return result, nil
}
