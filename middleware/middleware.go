package middleware

import (
	"net/http"

	"github.com/JulianDavidGamboa/basic-go-api/cmd/authorization"
	"github.com/labstack/echo/v4"
)

// Authentication .
func Authentication(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		_, err := authorization.ValidateToken(token)

		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "No permitido"})
		}

		return f(c)

	}
}
