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

func (controller *AuthController) Callback(c echo.Context) error {
	resp := model.AuthCallbackResponse{}

	req := model.AuthCallbackRequest{}
	if err := c.Bind(&req); err != nil {
		resp.Message = "不正なデータです。"
		return c.JSON(http.StatusBadRequest, &resp)
	}

	if err := validate.Struct(&req); err != nil {
		resp.Message = "不正なデータです。"
		return c.JSON(http.StatusBadRequest, &resp)
	}

	if err := model.VerifyState(c, req.State); err != nil {
		resp.Message = "stateが不正です。"
		return c.JSON(http.StatusBadRequest, &resp)
	}

	model.StoreCode(c, req.Code)
	return c.JSON(http.StatusOK, &resp)
}
