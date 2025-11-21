package service

import (
	"MIS/model"
)

func AddTeacher(t model.Teacher) (model.Teacher, error) {

	if err := ValidateTeacher(t); err != nil {
		return model.Teacher{}, err
	}

	return t, nil
}
