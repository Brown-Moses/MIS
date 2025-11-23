package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Person `GORM:"embedded"`
	DOB    string `json:"date_of_birth"`
	Grade  string `json:"grade"`
}
