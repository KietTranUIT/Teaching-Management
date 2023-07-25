package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase() *sql.DB {
	// Set up
	dbDriver := "mysql"
	dbUser := "kiettran"
	dbPassword := "Kiet@123456"
	dbName := "teachingManagement"

	// Connect to database
	db, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@tcp(127.0.0.1:3306)/"+dbName)

	if err != nil {
		log.Println("Connection to database failed")
		panic(err.Error())
	}

	return db
}

// Create a template object
var tmpl = template.Must(template.ParseFiles("index.html"))

// Handle when user access website
func Handle(w http.ResponseWriter, r *http.Request) {
	// Response a file index.html
	tmpl.Execute(w, nil)
}

func main() {
	// Connect to database on localhost
	db := ConnectDatabase()
	log.Println("Connection to database successful")
	fmt.Println(db)

	// Create a handler to process
	fs := http.FileServer(http.Dir("assets"))

	// Create a HTTP mux and register handle funcs
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", Handle)

	// Set up hostname, port, mux to run Server
	port := "8181"
	host := "127.0.0.1"
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: mux,
	}

	// Listen
	server.ListenAndServe()
}
