package validation

import (
	"MIS/model"
	"errors"
	"strings"
)

func ValidateStudent(s model.Student) error {

	// removes unnecessry space and returns error

	if strings.TrimSpace(s.FirstName) == "" {
		return errors.New("first name is required")
	}

	if strings.TrimSpace(s.LastName) == "" {
		return errors.New("last name is required")
	}

	if strings.TrimSpace(s.Gender) == "" {
		return errors.New("gender is required")
	}

	if strings.TrimSpace(s.DOB) == "" {
		return errors.New("date of birth is required")
	}

	return nil

}
