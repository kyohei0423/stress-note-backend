package router

import (
	"stress_note/infra/handler"

	"github.com/labstack/echo/v4"
)

func Router() {
	e := echo.New()
	e.POST("/signup", handler.Signup)
	e.Logger.Fatal(e.Start(":8080"))
}
