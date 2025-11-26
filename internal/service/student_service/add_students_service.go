package service

import (
	"MIS/internal/database"
	"MIS/internal/model"
	validate "MIS/internal/service/validation"
)

func AddStudent(s model.Student) (model.Student, error) {
	//validates using the validation function
	if err := validate.ValidateStudent(s); err != nil {
		return model.Student{}, err
	}

	//saves to database
	if err := database.DB.Create(&s).Error; err != nil {
		// returns to student
		return model.Student{}, err
	}

	return s, nil
}
