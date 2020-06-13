package view

import (
	"github.com/labstack/echo/v4"
	"github.com/mopeneko/donhaialert/api/model"
)

type SettingsPostView struct{}

func (v SettingsPostView) Render(c echo.Context, code int, message string) error {
	resp := model.SettingsPostResponse{Message: message}
	return c.JSON(code, &resp)
}

type SettingsDeleteView struct{}

func (v SettingsDeleteView) Render(c echo.Context, code int, message string) error {
	resp := model.SettingsDeleteResponse{Message: message}
	return c.JSON(code, &resp)
}
