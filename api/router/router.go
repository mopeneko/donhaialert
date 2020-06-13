package router

import (
	"crypto/rand"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/mopeneko/donhaialert/api/controller"
	"github.com/mopeneko/donhaialert/api/database"
	"github.com/mopeneko/donhaialert/api/domain"
)

var router *echo.Echo

func init() {
	router = echo.New()
}

func Init() {
	initSession()

	authGroup := router.Group("/auth")
	authController := &controller.AuthController{}
	authGroup.GET("", authController.Issue)
	authGroup.POST("/callback", authController.Callback)

	settingsGroup := router.Group("/settings")
	settingsController := &controller.SettingsController{}
	settingsGroup.POST("", settingsController.Enable)
	settingsGroup.DELETE("", settingsController.Disable)
}

func Run() {
	router.Logger.Fatal(router.Start(":1323"))
}

func initSession() {
	secret := domain.Secret{}
	err := database.DB.First(&secret).Error
	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			panic(fmt.Sprintf("failed to get secret: %+v", err))
		}

		genSecret, err := generateSecret()
		if err != nil {
			panic(fmt.Sprintf("failed to generate secret: %+v", err))
		}

		secret.Secret = genSecret
		database.DB.Create(&secret)
	}

	router.Use(session.Middleware(sessions.NewCookieStore(secret.Secret)))
}

func generateSecret() (b []byte, err error) {
	b = make([]byte, 64)
	_, err = rand.Read(b)
	return
}
