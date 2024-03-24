package data

import (
	"errors"
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
	Status        string    `gorm:"type:string" json:"status"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}

func (p *Plan) CreatePlan(userID uint, header, description string, timeDates ...time.Time) error {
	if timeDates[1].After(timeDates[2]) {
		return errors.New("start date cannot be after end date")
	}

	plan := Plan{
		UserID:        userID,
		Header:        header,
		Description:   description,
		Date:          timeDates[0],
		StartingHour:  timeDates[1],
		FinishingHour: timeDates[2],
		Status:        "Hazır",
	}

	return db.Create(&plan).Error
}

func (p *Plan) GetById(id string) (*Plan, error) {
	db.First(&p, "id = ?", id)
	return p, nil
}

func (p *Plan) UpdatePlan(header, description, status string, timeDates ...time.Time) *Plan {
	p.UpdatedAt = time.Now()
	p.Header = header
	p.Description = description
	p.Status = status
	p.Date = timeDates[0]
	p.StartingHour = timeDates[1]
	p.FinishingHour = timeDates[2]
	db.Save(&p)
	return p
}
