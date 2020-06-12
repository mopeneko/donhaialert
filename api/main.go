package main

import (
	"github.com/mopeneko/donhaialert/api/database"
	"github.com/mopeneko/donhaialert/api/router"
)

func main() {
	database.Init()
	database.DB.LogMode(true)
	router.Init()
	router.Run()
}
