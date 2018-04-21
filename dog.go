package main

import "github.com/jinzhu/gorm"

type Dog struct {
	gorm.Model
	Name     string `json:"name,omitempty"`
	Breed    string `json:"breed,omitempty"`
	ClientID int    `json:"oid,omitempty"`
}

func InsertDogs(dogs []Dog) int64 {
	return DBConn.Create(dogs).RowsAffected
}
