package models

import "gorm.io/gorm"

type Gender struct {
	gorm.Model
	Gender	string
}