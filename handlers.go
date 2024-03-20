package main

import (
	"github.com/labstack/echo/v4"
)

type JsonResponse struct {
	Error   bool
	Message string
	Data    any
}

func (app *Config) registerStudent(c echo.Context) error {
	var name, surname, email, password string
	name = c.FormValue("name")
	surname = c.FormValue("surname")
	email = c.FormValue("email")
	password = c.FormValue("password")

	err := app.Models.User.CreateUser(name, surname, email, password)
	if err != nil {
		return c.JSON(201, JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(201, JsonResponse{
		Error:   false,
		Message: "Student created succesfully",
		Data:    nil,
	})
}

func (app *Config) login(c echo.Context) error {
	var email, password string
	email = c.FormValue("email")
	password = c.FormValue("password")

	user, err := app.Models.User.GetByEmail(email)
	if err != nil {
		return c.JSON(400, JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	valid, err := user.PasswordMatches(password)
	if err != nil || !valid {
		return c.JSON(400, JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	token, err := GenerateToken(user.Email)
	if err != nil {
		return c.JSON(500, JsonResponse{
			Error:   true,
			Message: "Failed to generate JWT token",
			Data:    nil,
		})
	}

	return c.JSON(200, JsonResponse{
		Error:   false,
		Message: "Login successfully",
		Data:    token,
	})
}
