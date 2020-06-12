package domain

import "github.com/jinzhu/gorm"

type Credential struct {
	gorm.Model
	Host         string
	ClientID     string
	ClientSecret string
}
