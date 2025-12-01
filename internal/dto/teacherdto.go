package dto

type TeacherDTO struct {
	ID      uint   `json:"id,omitempty"`
	Name    string `json:"name" validate:"required"`
	Gender  string `json:"gender" validate:"required"`
	Subject string `json:"subject" validate:"required"`
}
