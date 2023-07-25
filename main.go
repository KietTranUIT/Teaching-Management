package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase() *sql.DB {
	dbDriver := "mysql"
	dbUser := "kiettran"
	dbPassword := "Kiet@123456"
	dbName := "teachingManagement"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@tcp(127.0.0.1:3306)/"+dbName)

	if err != nil {
		log.Println("Connection to database failed")
		panic(err.Error())
	}

	return db
}

func main() {
	db := ConnectDatabase()

	log.Println("Connection to database successful")
}
