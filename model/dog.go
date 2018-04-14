package model

import (
	"udemyTraining/http/BasicServer/util"
)

type Dog struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Breed string `json:"breed,omitempty"`
	Oid   int64  `json:"oid,omitempty"`
}


func GetClientDogs(clientId int64) []Dog {
	var dogs []Dog
	var dog Dog

	rows, err := DB.Query(`SELECT * FROM dogs WHERE oid = ?`, clientId)
	util.Check(err)

	for rows.Next() {
		err = rows.Scan(&dog.Id, &dog.Name, &dog.Breed, &dog.Oid)
		util.Check(err)
		dogs = append(dogs, dog)
	}

	return dogs
}

func InsertDogsForNewClient(dogs []Dog, clientEmail string) int64 {
	var numRowsInserted int64
	for _, dog := range dogs {
		stmt, err := DB.Prepare(`INSERT INTO dogs(name, breed, oid) SELECT ?, ?, id FROM clients WHERE email=?`)
		util.Check(err)

		result, err := stmt.Exec(dog.Name, dog.Breed, clientEmail)
		stmt.Close()
		util.Check(err)

		num, err := result.RowsAffected()
		numRowsInserted += num
	}

	return numRowsInserted
}

func InsertDogsForExistingClient(dogs []Dog) int64 {
	var numRecords int64
	for _, dog := range dogs {
		stmt, err := DB.Prepare(`INSERT INTO dogs(name, breed, oid) VALUES ($1, $2, $3)`)
		util.Check(err)

		r, err := stmt.Exec(dog.Name, dog.Breed, dog.Oid)
		stmt.Close()
		util.Check(err)

		n, err := r.RowsAffected()
		util.Check(err)

		numRecords += n
	}

	return numRecords
}
