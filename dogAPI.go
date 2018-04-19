package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func AddDogs(w http.ResponseWriter, req *http.Request) {
	var dogs []Dog
	_ = json.NewDecoder(req.Body).Decode(&dogs)
	numRowsInserted := InsertDogs(dogs)
	fmt.Fprintln(w, "SUCCESSFULLY INSERTED RECORDS: ", numRowsInserted)
}
