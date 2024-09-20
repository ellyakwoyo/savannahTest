package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"savannahTest/authentication"
)

func AuthenticateMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			accessToken := c.Request().Header.Get("Authorization")

			valid, err := authentication.ValidateSessionToken(accessToken[7:])
			if err != nil || !valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid session"})
			}

			return next(c)
		}
	}
}
