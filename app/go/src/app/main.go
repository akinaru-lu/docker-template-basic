package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	port = os.Getenv("PORT")

	dbName     = os.Getenv("DB_NAME")
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
)

func NewDBEngine() *sql.DB {
	src := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	fmt.Println("Connecting MySQL " + src)

	db, err := sql.Open("mysql", src)
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("MySQL connected")

	return db
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GO, "+r.RemoteAddr)
}

func main() {
	db := NewDBEngine()
	defer db.Close()

	http.HandleFunc("/", handler)

	fmt.Println("Serving :" + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err.Error())
	}
}
