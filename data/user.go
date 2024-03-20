package data

import (
	"errors"
	"net/mail"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(25)" json:"name"`
	Surname   string `gorm:"type:varchar(25)" json:"surname"`
	Email     string `gorm:"type:varchar(50);unique" json:"email"`
	Password  string `gorm:"type:varchar(255)"  json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) CreateUser(name, surname, email, password string) error {
	if name == "" || surname == "" {
		return errors.New("you need to enter name and surname")
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("you need to enter a valid email")
	}
	if len(password) < 5 {
		return errors.New("password must be minimum 5 length")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	user := User{
		Name:     name,
		Surname:  surname,
		Email:    email,
		Password: string(hashedPassword),
	}

	return db.Create(&user).Error
}

func (u *User) GetByEmail(email string) (*User, error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return &User{}, errors.New("you need to enter a valid")
	}

	db.First(&u, "email = ?", email)
	return u, nil
}

func (u *User) ResetPassword(password string) error {
	user, err := u.GetByEmail(u.Email)
	if err != nil {
		return err
	}

	if len(password) < 5 {
		return errors.New("password must be minimum 5 length")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	db.Save(&user)
	return nil
}

func (u *User) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	if err != nil {
		return false, err
	}
	return true, nil
}
