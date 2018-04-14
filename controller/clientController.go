package controller

import (
	"net/http"
	"encoding/json"
	"udemyTraining/http/BasicServer/model"
	"github.com/gorilla/mux"
	"strconv"
	"udemyTraining/http/BasicServer/util"
	"fmt"
)

func GetClients(w http.ResponseWriter, _ *http.Request) {
	var clients []model.Client
	clients = model.SelectClients()
	json.NewEncoder(w).Encode(clients)
}

func GetClient(w http.ResponseWriter, req *http.Request) {
	var client model.Client
	params := mux.Vars(req)
	clientId, err := strconv.ParseInt(params["id"], 10, 64)
	util.Check(err)
	client = model.SelectClient(clientId)
	json.NewEncoder(w).Encode(client)
}

func CreateClient(w http.ResponseWriter, req *http.Request) {
	var client model.Client
	_ = json.NewDecoder(req.Body).Decode(&client)
	numClientRecords, numDogRecords := model.InsertClient(client)
	fmt.Fprintln(w, "INSERTED CLIENT RECORDS", numClientRecords)
	fmt.Fprintln(w, "INSERTED DOG RECORDS", numDogRecords)
}
