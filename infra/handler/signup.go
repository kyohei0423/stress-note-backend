package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type user struct{}

func Signup(c echo.Context) error {
	res := user{}
	return c.JSON(http.StatusCreated, res)
}
