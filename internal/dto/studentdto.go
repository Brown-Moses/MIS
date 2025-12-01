package dto

type StudentDTO struct {
	ID        uint   `json:"id,omitempty"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Gender    string `json:"gender" validate:"required"`
	DOB       string `json:"date_of_birth" validate:"required"`
}
