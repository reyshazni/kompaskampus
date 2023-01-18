package lecture

import (
	"FindMyDosen/database"
	"FindMyDosen/model/entity"
)

func addLecture(lectureDTO *AddLectureDTO) error {
	db := database.GetDB()
	lecture := entity.Lecture{
		FullName:     lectureDTO.FullName,
		UniversityID: lectureDTO.UniversityID,
		Faculty:      lectureDTO.Faculty,
		Href:         lectureDTO.Href,
	}
	return db.Create(&lecture).Error
}

func getLectureById(id uint) (LectureDTO, error) {
	db := database.GetDB()
	lecture := new(entity.Lecture)
	//err := db.Preload("University").First(&lecture, "id = ?", id).Error
	err := db.Joins("lEFT JOIN universities on lectures.university_id = universities.id").First(&lecture, "id = ?", id).Error
	dto := LectureDTO{
		Id:       lecture.ID,
		FullName: lecture.FullName,
		//University: lecture.University.FullName,
		Faculty: lecture.Faculty,
		Href:    lecture.Href,
	}
	return dto, err
}

func getLecture(query *GetLectureQuery) ([]LectureDTO, error) {
	db := database.Paginate(query.Page, query.Limit)
	var lectures []entity.Lecture
	err := db.Joins("University").Where("lectures.full_name LIKE ? OR lectures.faculty LIKE ? OR University.full_name LIKE ? OR University.uni_code LIKE ? ", "%"+query.Search+"%", "%"+query.Search+"%", "%"+query.Search+"%", "%"+query.Search+"%").Find(&lectures).Error
	var result []LectureDTO
	for _, lecture := range lectures {
		result = append(result, LectureDTO{
			Id:       lecture.ID,
			FullName: lecture.FullName,
			//University: lecture.University.FullName,
			Faculty: lecture.Faculty,
			Href:    lecture.Href,
		})
	}
	return result, err
}
