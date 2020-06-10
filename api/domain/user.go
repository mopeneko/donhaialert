package domain

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	TootsCount      uint
	FollowsCount    uint
	FollowersSecret uint
}
