package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mopeneko/donhaialert/api/model"
	"net/http"
)

type AuthController struct {
}

func (controller *AuthController) Issue(c echo.Context) error {
	resp := model.AuthIssueResponse{}

	req := model.AuthIssueRequest{}
	req.Host = c.QueryParam("host")
	if err := validate.Struct(&req); err != nil {
		resp.Message = "ホストが正しくありません。"
		return c.JSON(http.StatusBadRequest, &resp)
	}

	credential, err := model.GetCredential(req.Host)
	if err != nil {
		resp.Message = "Credentialの取得に失敗しました。"
		return c.JSON(http.StatusInternalServerError, &resp)
	}

	url := model.GetAuthorizationURL(c, &credential)
	resp.Message = url
	return c.JSON(http.StatusOK, &resp)
}
