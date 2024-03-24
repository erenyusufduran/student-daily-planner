package data

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Plan struct {
	gorm.Model
	ID            uint      `gorm:"primaryKey" json:"id"`
	UserID        uint      `gorm:"index" json:"user_id"`
	Header        string    `gorm:"type:varchar(127)" json:"header"`
	Description   string    `gorm:"type:text" json:"description"`
	Date          time.Time `json:"date"`
	StartingHour  time.Time `json:"starting_hour"`
	FinishingHour time.Time `json:"finishing_hour"`
	Status        int       `gorm:"type:int" json:"status"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}

func (p *Plan) CreatePlan(userID uint, header, description string, date, startingHour, finishingHour time.Time) error {
	fmt.Println(userID)

	if startingHour.After(finishingHour) {
		return errors.New("start date cannot be after end date")
	}

	plan := Plan{
		UserID:        userID,
		Header:        header,
		Description:   description,
		Date:          date,
		StartingHour:  startingHour,
		FinishingHour: finishingHour,
		Status:        0,
	}

	return db.Create(&plan).Error
}
