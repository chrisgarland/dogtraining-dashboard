package main

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"udemyTraining/http/BasicServer/controller"
	"udemyTraining/http/BasicServer/model"
	"udemyTraining/http/BasicServer/util"
)

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

	model.DBInit()
	defer model.DB.Close()

	router := mux.NewRouter()

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/clients", controller.GetClients).Methods("GET")
	router.HandleFunc("/clients/", controller.GetClients).Methods("GET")
	router.HandleFunc("/clients/{id}", controller.GetClient).Methods("GET")
	router.HandleFunc("/new-client", controller.CreateClient).Methods("POST")
	router.HandleFunc("/new-dogs", controller.CreateDogs).Methods("POST")
	router.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "at index")
	util.Check(err)
}
