package controller

import (
	"net/http"
	"udemyTraining/http/BasicServer/model"
	"encoding/json"
	"fmt"
)

func CreateDogs(w http.ResponseWriter, req *http.Request)  {
	var dogs []model.Dog
	_ = json.NewDecoder(req.Body).Decode(&dogs)
	numRecords := model.InsertDogsForExistingClient(dogs)
	fmt.Fprintln(w, "INSERTED RECORD", numRecords)
}
