package main

import (
	db "Aurora02Api/database"
	"Aurora02Api/router"
)

var (
	host = "0.0.0.0:8099"
)

func main() {

	defer db.Eloquent.Close()

	router := router.InitRouter()

	router.Run(host)
}
