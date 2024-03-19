package main

import "github.com/labstack/echo/v4"

func (app *Config) registerStudent(c echo.Context) error {
	var name, surname, email, password string
	name = c.FormValue("name")
	surname = c.FormValue("surname")
	email = c.FormValue("email")
	password = c.FormValue("password")

	return app.Models.User.CreateUser(name, surname, email, password)
}
