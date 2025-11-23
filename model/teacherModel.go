package model

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Person  `GORM:"embedded"`
	Subject string `json:"subject"`
}
