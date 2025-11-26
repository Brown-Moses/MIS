package service

import (
	"MIS/internal/database"
	"MIS/internal/model"
	validate "MIS/internal/service/validation"
)

func AddTeacher(t model.Teacher) (model.Teacher, error) {

	//validates
	if err := validate.ValidateTeacher(t); err != nil {
		return model.Teacher{}, err
	}

	//saves to db
	if err := database.DB.Create(&t).Error; err != nil {
		// returns teachers
		return model.Teacher{}, err
	}

	return t, nil
}
