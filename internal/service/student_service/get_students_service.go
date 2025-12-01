package service

import (
	"MIS/internal/database"
	"MIS/internal/dto"
)

// get all student

func GetStudents() ([]dto.StudentDTO, error) {

	//temporary storage
	var student []dto.StudentDTO

	//load all members
	err := database.DB.Find(&student).Error
	if err != nil {
		return nil, err
	}

	return student, nil

}
