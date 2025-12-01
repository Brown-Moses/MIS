package service

import (
	"MIS/internal/database"
	"MIS/internal/dto"
	validate "MIS/internal/service/validation"
)

func AddTeacher(t dto.TeacherDTO) (dto.TeacherDTO, error) {

	//validates
	if err := validate.ValidateTeacher(t); err != nil {
		return dto.TeacherDTO{}, err
	}

	//saves to db
	if err := database.DB.Create(&t).Error; err != nil {
		// returns teachers
		return dto.TeacherDTO{}, err
	}

	return t, nil
}
