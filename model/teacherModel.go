package model

type Teacher struct {
	Person
	Subject string `json:"subject"`
}
