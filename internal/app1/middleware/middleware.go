package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
)

func MW(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("User-Role") == "admin" {
			log.Println("red button user detected")
		}

		err := next(c)

		return err
	}
}
