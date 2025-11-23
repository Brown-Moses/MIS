package service

import (
	"MIS/database"
	"MIS/model"
)

func AddTeacher(t model.Teacher) (model.Teacher, error) {

	//validates
	if err := ValidateTeacher(t); err != nil {
		return model.Teacher{}, err
	}

	//saves to db
	if err := database.DB.Create(&t).Error; err != nil {
		// returns teachers
		return model.Teacher{}, err
	}

	return t, nil
}
