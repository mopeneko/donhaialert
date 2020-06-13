package domain

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	TootsCount     int64
	FollowsCount   int64
	FollowersCount int64
	AccessTokenID  uint
	AccessToken    AccessToken
}
