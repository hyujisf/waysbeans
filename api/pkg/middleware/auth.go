package middleware

import (
	"net/http"
	"strings"
	"waysbeans/dto"
	jwtToken "waysbeans/pkg/jwt"

	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// mengambil token
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: "error", Message: "unauthorized"})
		}

		//memisahkan token dengan bearer token
		token = strings.Split(token, " ")[1]

		// validasi token dan mengambil claims
		claims, err := jwtToken.DecodeToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: "error", Message: "unathorized"})
		}

		// mengirim claims melalui context
		c.Set("userLogin", claims)
		return next(c)
	}
}

func AdminAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// mengambil token
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: "error", Message: "unauthorized"})
		}

		//memisahkan token dengan bearer token
		token = strings.Split(token, " ")[1]

		// validasi token dan mengambil claims
		claims, err := jwtToken.DecodeToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: "error", Message: "unathorized"})
		}

		if userRole := claims["role"].(string); userRole != "admin" {

			return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: "error", Message: "Unauthorized, You're not administrator"})
		}

		// mengirim claims melalui context
		c.Set("userLogin", claims)
		return next(c)
	}
}
