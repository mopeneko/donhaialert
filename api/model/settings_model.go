package model

import (
	"context"
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/mattn/go-mastodon"
	"github.com/mopeneko/donhaialert/api/database"
	"github.com/mopeneko/donhaialert/api/domain"
)

type SettingsPostResponse struct {
	Message string `json:"message"`
}

type SettingsDeleteResponse struct {
	Message string `json:"message"`
}

func SettingsEnable(c echo.Context) error {
	ctx := context.Background()

	sess, _ := session.Get("session", c)
	code := sess.Values["code"].(string)
	host := sess.Values["host"].(string)

	credential := domain.Credential{Host: host}
	database.DB.Where(&credential).First(&credential)

	config := mastodon.Config{
		Server:       "https://" + host,
		ClientID:     credential.ClientID,
		ClientSecret: credential.ClientSecret,
	}
	client := mastodon.NewClient(&config)
	err := client.AuthenticateToken(ctx, code, "https://donhaialert.com/callback")
	if err != nil {
		return errors.New("クライアントの生成に失敗しました。")
	}

	account, err := client.GetAccountCurrentUser(ctx)
	if err != nil {
		return errors.New("アカウント情報の取得に失敗しました。")
	}

	err = database.DB.
		Where(&domain.User{MastodonID: account.Username, CredentialID: credential.ID}).
		Find(&domain.User{}).
		Error
	if err == nil {
		return errors.New("既に登録されています。")
	}

	accessToken := domain.AccessToken{
		AccessToken: config.AccessToken,
	}

	user := domain.User{
		MastodonID:     account.Username,
		TootsCount:     account.StatusesCount,
		FollowsCount:   account.FollowingCount,
		FollowersCount: account.FollowersCount,
		AccessToken:    accessToken,
		Credential:     credential,
	}

	err = database.DB.Create(&user).Error
	if err != nil {
		return errors.New("登録に失敗しました。")
	}

	return nil
}

func SettingsDisable(c echo.Context) error {
	ctx := context.Background()

	sess, _ := session.Get("session", c)
	code := sess.Values["code"].(string)
	host := sess.Values["host"].(string)

	credential := domain.Credential{Host: host}
	database.DB.Where(&credential).First(&credential)

	config := mastodon.Config{
		Server:       "https://" + host,
		ClientID:     credential.ClientID,
		ClientSecret: credential.ClientSecret,
	}
	client := mastodon.NewClient(&config)
	err := client.AuthenticateToken(ctx, code, "https://donhaialert.com/callback")
	if err != nil {
		return errors.New("クライアントの生成に失敗しました。")
	}

	account, err := client.GetAccountCurrentUser(ctx)
	if err != nil {
		return errors.New("アカウント情報の取得に失敗しました。")
	}

	user := &domain.User{
		MastodonID:   account.Username,
		CredentialID: credential.ID,
	}

	err = database.DB.
		Where(&user).
		First(&user).
		Find(&user).
		Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("既に削除されています。")
		}
		log.Printf("[Error] [SettingsDisable] %+v\n", err)
		return errors.New("予期しないエラーが発生しました。")
	}

	err = database.DB.Delete(&user).Error
	if err != nil {
		return errors.New("ユーザー情報の削除に失敗しました。")
	}

	err = database.DB.Delete(&user.AccessToken).Error
	if err != nil {
		return errors.New("アクセストークンの削除に失敗しました。")
	}

	return nil
}
