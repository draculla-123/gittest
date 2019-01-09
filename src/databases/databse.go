package main

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
)

const (
	url = "localhost"
)

func main() {
	// connecting to mongodb server
	// ----
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	log.Printf("Successfully connected to mongodb server at %v", url)

	// Show all database names available
	dbNames, err := session.DatabaseNames()
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range dbNames {
		fmt.Printf("[%v] - %v\n", i+1, v)
	}
}
