package service

import (
	"MIS/database"
	"MIS/model"
)

// get all students by ID
func GetStudentsByID(id uint) (model.Student, error) {

	//first thing first, we declare in temporary memo
	var student model.Student

	//get with id
	err := database.DB.First(&student, id).Error
	if err != nil {
		return model.Student{}, err
	}

	return student, nil

}
