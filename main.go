package main

import (
	"Management/User"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"crypto/sha256"
	"encoding/hex"

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
		_, err = db.Exec(fmt.Sprintf("INSERT INTO User VALUES('%s', '%s', '%s', '%s', '%s')", user.GetUsername(), user.GetPassword(), user.GetId(), user.GetRole(), user.GetEmail()))
	case User.Student:
		student := value
		_, err = db.Exec(fmt.Sprintf("INSERT INTO SINHVIEN VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s')", student.GetMahv(), student.GetHo(), student.GetTen(), student.GetNgsinh(), student.GetGioitinh(), student.GetNoisinh(), student.GetMalop()))
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

		user.SetUser(query.Get("username"), query.Get("password"), "001", "no", query.Get("email"))
		fmt.Println(user)
		student.SetStudent("001", query.Get("ho"), query.Get("ten"), query.Get("ngaysinh"), query.Get("gioitinh"), query.Get("noisinh"), "01")

		fmt.Println(query.Get("ho"))
		fmt.Println(student)
		ok := Insert(db, user)

		ok1 := Insert(db, student)

		if ok != nil {
			log.Println(ok.Error())
			return
		}
		if ok1 != nil {
			log.Println(ok1.Error())
			return
		}
		log.Println("Insert succeeded!")
	}
}

func HandleLoginForStudent(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, _ := url.Parse(r.URL.String())
		form := u.Query()

		username := form.Get("username")
		password := form.Get("password")

		hash := sha256.Sum256([]byte(password))
		password = hex.EncodeToString(hash[:])

		query := fmt.Sprintf("SELECT password FROM User WHERE username='%s'", username)

		row, _ := db.Query(query)
		var pass string

		for row.Next() {
			row.Scan(&pass)
		}

		if pass == password {
			http.ServeFile(w, r, "main.html")
		}
	}
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
	mux.HandleFunc("/Student/Login/login-student", HandleLoginForStudent(db))
	mux.HandleFunc("/Student/Register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "registerStudent.html")
	})
	mux.HandleFunc("/Student/Register/register-student", HandleRegisterForStudent(db))
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
