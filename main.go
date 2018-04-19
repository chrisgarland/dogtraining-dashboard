package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"os"
)

var (
	DBConn *gorm.DB
	err    error
)

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/api/clients", GetClients).Methods("GET")
	router.HandleFunc("/api/clients/{id}", GetClient).Methods("GET")
	router.HandleFunc("/api/clients", CreateClient).Methods("POST")
	router.HandleFunc("/api/clients/batch", CreateClients).Methods("POST")
	router.HandleFunc("/api/dogs", AddDogs).Methods("POST")
	router.Handle("/favicon.ico", http.NotFoundHandler())
    log.Fatal(http.ListenAndServe(":80", router))
}

func dbInit() {
	DBConn.CreateTable(&Client{})
	DBConn.CreateTable(&Dog{})
	DBConn.AutoMigrate(&Client{})
	DBConn.AutoMigrate(&Dog{})
}

func initLog() *os.File {
	f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("ERROR:", err)
	} else {
		log.SetOutput(f)
		log.Println(`Log output successfully set to "log.txt"`)
	}

	return f
}

func main() {
	logfile := initLog()
	defer logfile.Close()

	dataSourceName := "administrator:administrator@tcp(testdb.c11gtkstiblb.ap-southeast-2.rds.amazonaws.com:3306)/test?charset=utf8&parseTime=True&loc=Local"
	DBConn, err = gorm.Open("mysql", dataSourceName)
	defer DBConn.Close()
	CheckFatal(err)
	dbInit()

	handleRequests()
}
