package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mopeneko/donhaialert/api/model"
	"github.com/mopeneko/donhaialert/api/view"
	"net/http"
)

type SettingsController struct{}

func (controller *SettingsController) Enable(c echo.Context) error {
	v := view.SettingsPostView{}

	err := model.SettingsEnable(c)
	if err != nil {
		return v.Render(c, http.StatusInternalServerError, err.Error())
	}

	return v.Render(c, http.StatusOK, "")
}

func (controller *SettingsController) Disable(c echo.Context) error {
	v := view.SettingsDeleteView{}

	err := model.SettingsDisable(c)
	if err != nil {
		return v.Render(c, http.StatusInternalServerError, err.Error())
	}

	return v.Render(c, http.StatusOK, "")
}
