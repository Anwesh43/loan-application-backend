package models

import "gorm.io/gorm"

type Application struct {
	gorm.Model
	Id              int `gorm:"primaryKey;autoIncrement"`
	Name            string
	YearEstablished string
	Status          string
}
