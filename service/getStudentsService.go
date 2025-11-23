package service

import (
	"MIS/database"
	"MIS/model"
)

// get all student

func GetStudents() ([]model.Student, error) {

	//temporary storage
	var student []model.Student

	//load all members
	err := database.DB.Find(&student).Error
	if err != nil {
		return nil, err
	}

	return student, nil

}
