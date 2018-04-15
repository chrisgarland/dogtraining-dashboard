package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var db *sql.DB
var err error

type Client struct {
	Id        int64  `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Dogs      []Dog  `json:"dogs,omitempty"`
}

type Dog struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Breed string `json:"breed,omitempty"`
	Oid   int64  `json:"oid,omitempty"`
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "at index")
	check(err)
}

func getClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var client Client
	var clients []Client

	ownerRows, err := db.Query(`SELECT * FROM clients;`)
	check(err)
	defer ownerRows.Close()

	for ownerRows.Next() {
		err = ownerRows.Scan(&client.Id, &client.Email, &client.Firstname, &client.Lastname)
		check(err)

		client.Dogs = getClientDogs(client.Id)
		clients = append(clients, client)
	}

	json.NewEncoder(w).Encode(clients)
}

func getClient(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var client Client
	params := mux.Vars(req)

	clientId, err := strconv.ParseInt(params["id"], 10, 64)
	check(err)

	row := db.QueryRow(`SELECT * FROM clients WHERE id=?`, clientId)

	err = row.Scan(&client.Id, &client.Email, &client.Firstname, &client.Lastname)
	check(err)

	client.Dogs = getClientDogs(client.Id)

	json.NewEncoder(w).Encode(client)
}

func getClientDogs(clientId int64) []Dog {
	var dogs []Dog
	var dog Dog

	rows, err := db.Query(`SELECT * FROM dogs WHERE oid = ?`, clientId)
	check(err)

	for rows.Next() {
		err = rows.Scan(&dog.Id, &dog.Name, &dog.Breed, &dog.Oid)
		check(err)
		dogs = append(dogs, dog)
	}

	return dogs
}

func insertClient(w http.ResponseWriter, req *http.Request) {
	var client Client
	_ = json.NewDecoder(req.Body).Decode(&client)
	stmt, err := db.Prepare(`INSERT INTO clients(email, fname, lname) VALUES (?, ?, ?)`)

	check(err)
	defer stmt.Close()

	r, err := stmt.Exec(client.Email, client.Firstname, client.Lastname)
	check(err)

	n, err := r.RowsAffected()
	check(err)

	numDogRecords := insertDogsForNewClient(client.Dogs, client.Email)

	fmt.Fprintln(w, "INSERTED CLIENT RECORDS", n)
	fmt.Fprintln(w, "INSERTED DOG RECORDS", numDogRecords)
}

func insertDogsForNewClient(dogs []Dog, clientEmail string) int64 {
	var numRowsInserted int64
	for _, dog := range dogs {
		stmt, err := db.Prepare(`INSERT INTO dogs(name, breed, oid) SELECT ?, ?, id FROM clients WHERE email=?`)
		check(err)

		result, err := stmt.Exec(dog.Name, dog.Breed, clientEmail)
		stmt.Close()
		check(err)

		num, err := result.RowsAffected()
		numRowsInserted += num
	}

	return numRowsInserted
}

func insertDogsForExistingClient(w http.ResponseWriter, req *http.Request) {
	var dogs []Dog
	_ = json.NewDecoder(req.Body).Decode(&dogs)

	for _, dog := range dogs {
		stmt, err := db.Prepare(`INSERT INTO dogs(name, breed, oid) VALUES ($1, $2, $3)`)
		check(err)

		r, err := stmt.Exec(dog.Name, dog.Breed, dog.Oid)
		stmt.Close()
		check(err)

		n, err := r.RowsAffected()
		check(err)

		fmt.Fprintln(w, "INSERTED RECORD", n)
	}
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
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

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/clients", getClients).Methods("GET")
	router.HandleFunc("/clients/{id}", getClient).Methods("GET")
	router.HandleFunc("/clients", insertClient).Methods("POST")
	router.HandleFunc("/dogs", insertDogsForExistingClient).Methods("POST")
	router.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	logfile := initLog()
	defer logfile.Close()

	dataSourceName := "administrator:administrator@tcp(testdb.c11gtkstiblb.ap-southeast-2.rds.amazonaws.com:3306)/test?charset=utf8"
	if db, err = sql.Open("mysql", dataSourceName); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	check(err)

	handleRequests()
}
