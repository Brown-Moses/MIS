package service

import (
	"MIS/internal/database"
	"MIS/internal/dto"
	validate "MIS/internal/service/validation"
)

func AddStudent(s dto.StudentDTO) (dto.StudentDTO, error) {
	//validates using the validation function
	if err := validate.ValidateStudent(s); err != nil {
		return dto.StudentDTO{}, err
	}

	//saves to database
	if err := database.DB.Create(&s).Error; err != nil {
		// returns to student
		return dto.StudentDTO{}, err
	}

	return s, nil
}
