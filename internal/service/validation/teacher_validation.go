package validation

import (
	"MIS/internal/model"
	"errors"
	"strings"
)

func ValidateTeacher(t model.Teacher) error {

	if strings.TrimSpace(t.FirstName) == "" {
		return errors.New("first name is required")
	}

	if strings.TrimSpace(t.LastName) == "" {
		return errors.New("last name is required")
	}

	if strings.TrimSpace(t.Gender) == "" {
		return errors.New("gender is required")
	}

	return nil
}
