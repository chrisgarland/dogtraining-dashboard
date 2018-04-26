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

func SelectDos() []Dog {
	var dogs []Dog
	DBConn.Find(&dogs)
	return dogs
}

func DeleteDogs() error {
	sqlString := "DELETE FROM dogs WHERE oid > ?"
	return DBConn.Exec(sqlString, 10).Error
}
