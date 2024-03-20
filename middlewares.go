package main

import (
	"errors"

	"github.com/labstack/echo/v4"
)

func (app *Config) authMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			email, err := VerifyToken(token)
			if token == "" || err != nil {
				return errors.New("error occured")
			}

			c.Set("email", email)
			return next(c)
		}
	}
}
