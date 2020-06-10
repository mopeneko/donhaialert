package domain

import "github.com/jinzhu/gorm"

type Credential struct {
	gorm.Model
	Host         string
	ClientId     string
	ClientSecret string
}
