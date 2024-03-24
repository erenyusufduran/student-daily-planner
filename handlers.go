package main

import (
	"github.com/erenyusufduran/student-lesson/data"
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

	token, err := GenerateToken(user.Email, int64(user.ID))
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

func (app *Config) createPlan(c echo.Context) error {
	userId := c.Get("userId")
	var header, description, startingHour, finishingHour, date string
	header = c.FormValue("header")
	description = c.FormValue("description")
	startingHour = c.FormValue("startingHour")
	finishingHour = c.FormValue("finishingHour")
	date = c.FormValue("date")

	if header == "" || description == "" || startingHour == "" || finishingHour == "" || date == "" {
		return c.JSON(400, JsonResponse{
			Error:   true,
			Message: "All fields must be filled (header, description, startingHour, finishingHour, date)",
			Data:    nil,
		})
	}

	timeDates, err := timeWithDateStartingFinishing(date, startingHour, finishingHour)
	if err != nil {
		return c.JSON(400, JsonResponse{
			Error:   true,
			Message: "Date calculation failed",
			Data:    nil,
		})
	}

	err = app.Models.Plan.CreatePlan(uint(userId.(int64)), header, description, timeDates[0], timeDates[1], timeDates[2])
	if err != nil {
		return c.JSON(400, JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(201, JsonResponse{
		Error:   false,
		Message: "Plan created successfully",
		Data: data.Plan{
			Header:        header,
			Description:   description,
			Date:          timeDates[0],
			StartingHour:  timeDates[1],
			FinishingHour: timeDates[2],
		},
	})
}

func (app *Config) updatePlan(c echo.Context) error {
	id := c.Param("id")
	userId := c.Get("userId")
	plan, err := app.Models.Plan.GetById(id)
	if err != nil {
		return c.JSON(400, JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if plan.UserID != uint(userId.(int64)) {
		return c.JSON(400, JsonResponse{
			Error:   true,
			Message: "Unauthorized, this is not your plan",
			Data:    nil,
		})
	}

	var header, description, startingHour, finishingHour, date string
	header = c.FormValue("header")
	description = c.FormValue("description")
	startingHour = c.FormValue("startingHour")
	finishingHour = c.FormValue("finishingHour")
	date = c.FormValue("date")
	status := c.FormValue("status")

	if header == "" || description == "" || startingHour == "" || finishingHour == "" || date == "" || status == "" {
		return c.JSON(400, JsonResponse{
			Error:   true,
			Message: "All fields must be filled (header, description, startingHour, finishingHour, date)",
			Data:    nil,
		})
	}
	if status == "Hazır" || status == "İptal" || status == "Yapılıyor" || status == "Bitti" {
		timeDates, err := timeWithDateStartingFinishing(date, startingHour, finishingHour)
		if err != nil {
			return c.JSON(400, JsonResponse{
				Error:   true,
				Message: "Date calculation failed",
				Data:    nil,
			})
		}

		plan = plan.UpdatePlan(header, description, status, timeDates...)

		return c.JSON(200, JsonResponse{
			Error:   false,
			Message: "Plan updated",
			Data:    plan,
		})
	}

	return c.JSON(400, JsonResponse{
		Error:   false,
		Message: "Status musts be Hazır, İptal, Yapılıyor, Bitti",
		Data:    plan,
	})
}
