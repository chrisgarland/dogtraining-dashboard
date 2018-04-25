package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func GetClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var clientsDTO [][5]string
	clients := SelectClients()

	for _, client := range clients {
		clientDto := [5]string {
			fmt.Sprint(client.ID),
			client.Email,
			client.Firstname,
			client.Lastname,
			time.Time.String(client.UpdatedAt),
		}
		clientsDTO  = append(clientsDTO, clientDto)
	}

	json.NewEncoder(w).Encode(clientsDTO)
}

func GetClient(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	clientId, err := strconv.ParseInt(params["id"], 10, 64)
	Check(err)
	json.NewEncoder(w).Encode(SelectClient(clientId))
}

func CreateClients(w http.ResponseWriter, r *http.Request) {
	var clients []Client
	var numRecords int64
	_ = json.NewDecoder(r.Body).Decode(&clients)
	for _, client := range clients {
		numRecords += InsertClient(client)
	}
	fmt.Fprintln(w, "SUCCESSFULLY CREATED RECORDS: ", numRecords)
}

func CreateClient(w http.ResponseWriter, req *http.Request) {
	var client Client
	_ = json.NewDecoder(req.Body).Decode(&client)
	numRecordsCreated := InsertClient(client)
	fmt.Fprintln(w, "SUCCESSFULLY CREATED RECORDS: ", numRecordsCreated)
}
