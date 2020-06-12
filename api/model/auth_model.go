package model

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/mattn/go-mastodon"
	"github.com/mopeneko/donhaialert/api/database"
	"github.com/mopeneko/donhaialert/api/domain"
	"golang.org/x/oauth2"
	"log"
)

type AuthIssueRequest struct {
	Host string `validate:"required,fqdn"`
}

type AuthIssueResponse struct {
	Message string `json:"message"`
}

func GetCredential(host string) (credential domain.Credential, err error) {
	credential = domain.Credential{Host: host}
	err = database.DB.Where(&credential).First(&credential).Error
	if (err != nil && !gorm.IsRecordNotFoundError(err)) || len(credential.ClientID) > 0 {
		return
	}

	app, err := mastodon.RegisterApp(context.Background(), &mastodon.AppConfig{
		Server:       "https://" + host,
		ClientName:   "donhaialert",
		RedirectURIs: "https://api.donhaialert.com/auth/callback",
		Scopes:       "read:accounts write:statuses",
		Website:      "https://donhaialert.com/",
	})
	if err != nil {
		return
	}

	credential.ClientID = app.ClientID
	credential.ClientSecret = app.ClientSecret
	database.DB.Create(&credential)

	log.Println("Application generated:", host)
	return
}

func GetAuthorizationURL(c echo.Context, credential *domain.Credential) string {
	config := oauth2.Config{
		ClientID:     credential.ClientID,
		ClientSecret: credential.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://" + credential.Host + "/oauth/authorize",
			TokenURL:  "https://" + credential.Host + "/oauth/token",
			AuthStyle: oauth2.AuthStyleInParams,
		},
		RedirectURL: "https://api.donhaialert.com/auth/callback",
		Scopes:      []string{"read:accounts", "write:statuses"},
	}

	state := generateRandomState()

	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{}
	sess.Values["state"] = state
	sess.Save(c.Request(), c.Response())

	url := config.AuthCodeURL(state)
	return url
}

func generateRandomState() string {
	b := make([]byte, 64)
	rand.Read(b)
	state := base64.StdEncoding.EncodeToString(b)
	return state
}
