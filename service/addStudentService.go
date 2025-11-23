package service

import (
	"MIS/database"
	"MIS/model"
)

func AddStudent(s model.Student) (model.Student, error) {
	//validates using the validation function
	if err := ValidateStudent(s); err != nil {
		return model.Student{}, err
	}

	//saves to database
	if err := database.DB.Create(&s).Error; err != nil {
		// returns to student
		return model.Student{}, err
	}

	return s, nil
}
