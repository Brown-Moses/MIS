package service

import (
	"MIS/model"
)

func AddStudent(s model.Student) (model.Student, error) {
	if err := ValidateStudent(s); err != nil {
		return model.Student{}, err
	}
	return s, nil
}
