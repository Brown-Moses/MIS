package validation

import (
	"MIS/internal/dto"
	"errors"
	"strings"
)

func ValidateTeacher(t dto.TeacherDTO) error {

	if strings.TrimSpace(t.Name) == "" {
		return errors.New("name is required")
	}

	if strings.TrimSpace(t.Gender) == "" {
		return errors.New("gender is required")
	}

	return nil
}
