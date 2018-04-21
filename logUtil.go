package main

import "log"

func Check(err error) {
	if err != nil {
		log.Println(err)
	}
}

func CheckFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
