package models

import "gorm.io/gorm"

type ApplicationSummary struct {
	gorm.Model
	CurrApplication   Application `gorm:"foreignKey:Id"`
	PLInLastYear      int
	AverageAssetValue int
}
