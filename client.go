package main

import (
	"github.com/jinzhu/gorm"
)

type Client struct {
	gorm.Model
	Email     string `json:"email,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Dogs      []Dog  `json:"dogs,omitempty"`
}

func SelectClients() []Client {
	var clients []Client
	DBConn.Preload("Dogs").Find(&clients)
	return clients
}

func SelectClient(clientId int64) Client {
	var client Client
	DBConn.Preload("Dogs").First(&client, clientId)
	return client
}

func InsertClient(client Client) int64 {
	return DBConn.Create(&client).RowsAffected
}
