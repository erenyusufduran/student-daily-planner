package data

import (
	"time"

	"gorm.io/gorm"
)

type Plan struct {
	gorm.Model
	User          User   `gorm:"foreignKey:ID"`
	Header        string `gorm:"type:varchar(127)"`
	Description   string `gorm:"type:text"`
	Date          time.Time
	StartingHour  time.Time
	FinishingHour time.Time
	Status        string `gorm:"type:varchar(255)"`
}
