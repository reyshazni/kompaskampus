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
