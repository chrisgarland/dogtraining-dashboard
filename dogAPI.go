package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

func GetClientDogs(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	clientId, err := strconv.ParseInt(params["client_id"], 10, 64)
	Check(err)
	json.NewEncoder(w).Encode(SelectClientDogs(clientId))
}

func AddDogs(w http.ResponseWriter, req *http.Request) {
	var dogs []Dog
	_ = json.NewDecoder(req.Body).Decode(&dogs)
	numRowsInserted := InsertDogs(dogs)
	fmt.Fprintln(w, "SUCCESSFULLY INSERTED RECORDS: ", numRowsInserted)
}
