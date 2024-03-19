package data

import "gorm.io/gorm"

var db *gorm.DB

type Models struct {
	User User
	Plan Plan
}

func New(dbPool *gorm.DB) Models {
	db = dbPool

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Plan{})

	return Models{
		User: User{},
		Plan: Plan{},
	}
}
