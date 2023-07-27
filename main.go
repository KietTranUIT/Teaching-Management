package main

import (
	"Management/User"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

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

func Insert(db *sql.DB, data any) error {
	var err error
	switch value := data.(type) {
	case User.User:
		user := value
		_, err = db.Exec(fmt.Sprintf("INSERT INTO User VALUES('%s', '%s')", user.GetUsername(), user.GetPassword()))
	case User.Student:
		student := value
		_, err = db.Exec(fmt.Sprintf("INSERT INTO SINHVIEN VALUES('%s', '%s', '%s', '%s', '%s', '%s' '%s')", student.GetMahv(), student.GetHo(), student.GetTen(), student.GetNgsinh(), student.GetGioitinh(), student.GetNoisinh(), student.GetMalop()))
	}
	return err
}

// Create a template object
var tmpl = template.Must(template.ParseFiles("index.html"))

// Handle when user access website
func Handle(w http.ResponseWriter, r *http.Request) {
	// Response a file index.html
	tmpl.Execute(w, nil)
}

func HandleRegisterForStudent(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(r.URL.String())

		if err != nil {
			log.Println(err.Error())
		}

		query := u.Query()

		var user User.User
		var student User.Student

		user.SetUser(query.Get("username"), query.Get("password"))
		student.SetStudent("001", query.Get("ho"), query.Get("ten"), query.Get("ngaysinh"), query.Get("gioitinh"), query.Get("noisinh"), "")

		ok := Insert(db, user)

		Insert(db, student)

		if ok != nil {
			log.Println(ok.Error())
			return
		}

		log.Printf("Successfully add username: %s, password: %s to database\n", user.GetUsername(), user.GetPassword())
	}
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login request")
}

func main() {
	// Connect to database on localhost
	db := ConnectDatabase()
	log.Println("Connection to database successful")

	// Create a handler to process
	fs := http.FileServer(http.Dir("assets"))

	// Create a HTTP mux and register handle funcs
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/Student/Login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "student.html")
	})
	mux.HandleFunc("/Student/Login/login-student", HandleLogin)
	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ok")
	})
	mux.HandleFunc("/sign_up", HandleRegisterForStudent(db))
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
