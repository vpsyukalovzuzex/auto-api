package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func GetSub(c echo.Context) int {
	t := c.Get("user").(*jwt.Token)
	claims := t.Claims.(jwt.MapClaims)
	return int(claims["sub"].(float64))
}
