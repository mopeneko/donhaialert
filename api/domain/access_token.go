package domain

import "github.com/jinzhu/gorm"

type AccessToken struct {
	gorm.Model
	AccessToken  string
	CredentialID int
	Credential   Credential
	UserID       int
	User         User
}
