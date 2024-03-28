package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//type users []user

func main() {
	db, err := getDB()
	if err != nil {
		log.Fatal("1", err)
	}
	defer db.Close()

	router := getRouter(db)
	router.Run("0.0.0.0:9000")
}
