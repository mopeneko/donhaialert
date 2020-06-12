package view

import (
	"github.com/labstack/echo/v4"
	"github.com/mopeneko/donhaialert/api/model"
)

type AuthIssueView struct{}

func (v AuthIssueView) Render(c echo.Context, code int, message string) error {
	resp := model.AuthIssueResponse{Message: message}
	return c.JSON(code, &resp)
}

type AuthCallbackView struct{}

func (v AuthCallbackView) Render(c echo.Context, code int, message string) error {
	resp := model.AuthIssueResponse{Message: message}
	return c.JSON(code, &resp)
}
