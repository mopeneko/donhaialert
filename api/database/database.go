package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mopeneko/donhaialert/api/domain"
	"log"
	"net/url"
	"time"
)

const (
	host     = "mysql"
	user     = "mopeneko"
	password = "mopepass"
	dbname   = "donhaialert"
	loc      = "Asia/Tokyo"
)

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		user, password, host, dbname, true, url.QueryEscape(loc),
	)

	DB = connect(dsn)

	DB.AutoMigrate(&domain.Credential{}, &domain.User{}, &domain.AccessToken{}, &domain.Secret{})
}

func connect(dsn string) *gorm.DB {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Println("Failed to connect to database. Restarting...")
		log.Printf("Error: %+v\n", err)
		time.Sleep(time.Second * 3)
		return connect(dsn)
	}

	return db
}
