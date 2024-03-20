package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *Config) authMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			fmt.Print(token)
			if token == "" {
				return errors.New("Missing token")
			}

			valid := true

			if !valid {
				return c.JSON(http.StatusUnauthorized, "Invalid token")
			}

			// Token is valid, continue to the next handler
			return next(c)
		}
	}
}
