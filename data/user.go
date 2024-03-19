package data

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(25)"`
	Surname   string `gorm:"type:varchar(25)"`
	Email     string `gorm:"type:varchar(50);unique"`
	Password  string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) CreateUser(name, surname, email, password string) error {
	user := User{
		Name:     name,
		Surname:  surname,
		Email:    email,
		Password: password,
	}

	return db.Create(&user).Error
}
