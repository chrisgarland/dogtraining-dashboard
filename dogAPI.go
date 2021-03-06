package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetDogs(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(SelectDos())
}

func AddDogs(w http.ResponseWriter, req *http.Request) {
	var dogs []Dog
	_ = json.NewDecoder(req.Body).Decode(&dogs)
	numRowsInserted := InsertDogs(dogs)
	fmt.Fprintln(w, "SUCCESSFULLY INSERTED RECORDS: ", numRowsInserted)
}

func RemoveDogs(w http.ResponseWriter, _ *http.Request) {
	if err := DeleteDogs(); err != nil {
		fmt.Fprintln(w, err)
	} else {
		fmt.Fprintln(w, "Successfully deleted dogs")
	}
}
