package domain

import "github.com/jinzhu/gorm"

type Secret struct {
	gorm.Model
	Secret []byte
}
