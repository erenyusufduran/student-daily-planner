package main

import (
	"github.com/labstack/echo/v4"
)

func (app *Config) authMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			userId, err := VerifyToken(token)
			if token == "" || err != nil {
				return c.JSON(400, JsonResponse{
					Error:   true,
					Message: "Token could not verified",
					Data:    nil,
				})
			}

			c.Set("userId", userId)
			return next(c)
		}
	}
}
