package view

import "github.com/labstack/echo/v4"

type View interface {
	Render(echo.Context, int, string) error
}
