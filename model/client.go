package model

import (
	"udemyTraining/http/BasicServer/util"
)

type Client struct {
	Id        int64  `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Dogs      []Dog  `json:"dogs,omitempty"`
}

func SelectClients() []Client {
	var client Client
	var clients []Client

	ownerRows, err := DB.Query(`SELECT * FROM clients;`)
	util.Check(err)
	defer ownerRows.Close()

	for ownerRows.Next() {
		err = ownerRows.Scan(&client.Id, &client.Email, &client.Firstname, &client.Lastname)
		util.Check(err)

		client.Dogs = GetClientDogs(client.Id)
		clients = append(clients, client)
	}

	return clients
}

func SelectClient(clientId int64) Client {
	var client Client

	row := DB.QueryRow(`SELECT * FROM clients WHERE id=?`, clientId)

	err := row.Scan(&client.Id, &client.Email, &client.Firstname, &client.Lastname)
	util.Check(err)

	client.Dogs = GetClientDogs(client.Id)

	return client
}

func InsertClient(client Client) (int64, int64) {
	stmt, err := DB.Prepare(`INSERT INTO clients(email, fname, lname) VALUES (?, ?, ?)`)
	util.Check(err)
	defer stmt.Close()

	r, err := stmt.Exec(client.Email, client.Firstname, client.Lastname)
	util.Check(err)

	nunClientRecords, err := r.RowsAffected()
	util.Check(err)

	numDogRecords := InsertDogsForNewClient(client.Dogs, client.Email)

	return nunClientRecords, numDogRecords
}
