package domain

import "github.com/jinzhu/gorm"

type AccessToken struct {
	gorm.Model
	AccessToken  string
	CredentialID uint
	Credential   Credential
}
