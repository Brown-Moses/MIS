package service

import (
	"MIS/internal/database"
	"MIS/internal/dto"
)

// get all students by ID
func GetStudentsByID(id uint) (dto.StudentDTO, error) {

	//first thing first, we declare in temporary memo
	var student dto.StudentDTO

	//get with id
	err := database.DB.First(&student, id).Error
	if err != nil {
		return dto.StudentDTO{}, err
	}

	return student, nil

}
