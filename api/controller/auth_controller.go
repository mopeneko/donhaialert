package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mopeneko/donhaialert/api/model"
	"github.com/mopeneko/donhaialert/api/view"
	"net/http"
)

type AuthController struct {
}

func (controller *AuthController) Issue(c echo.Context) error {
	v := view.AuthIssueView{}

	req := model.AuthIssueRequest{}
	req.Host = c.QueryParam("host")
	if err := validate.Struct(&req); err != nil {
		return v.Render(c, http.StatusBadRequest, "ホストが正しくありません。")
	}

	credential, err := model.GetCredential(req.Host)
	if err != nil {
		return v.Render(c, http.StatusInternalServerError, "Credentialの取得に失敗しました。")
	}

	url := model.GetAuthorizationURL(c, &credential)
	return v.Render(c, http.StatusOK, url)
}

func (controller *AuthController) Callback(c echo.Context) error {
	v := view.AuthCallbackView{}

	req := model.AuthCallbackRequest{}
	if err := c.Bind(&req); err != nil {
		return v.Render(c, http.StatusBadRequest, "不正なデータです。")
	}

	if err := validate.Struct(&req); err != nil {
		return v.Render(c, http.StatusBadRequest, "不正なデータです。")
	}

	if err := model.VerifyState(c, req.State); err != nil {
		return v.Render(c, http.StatusBadRequest, "stateが不正です。")
	}

	model.StoreCode(c, req.Code)
	return v.Render(c, http.StatusOK, "")
}
