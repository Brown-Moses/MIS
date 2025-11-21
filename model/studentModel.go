package model

type Student struct {
	Person
	DOB   string `json:"date_of_birth"`
	Grade string `json:"grade"`
}
