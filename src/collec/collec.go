package main

import (
	"fmt"
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

const (
	url = "localhost"
)

func main() {
	dbName := "test"
	if 1 == len(os.Args) {
		log.Printf("No db specified, using '%v'", dbName)
	} else {
		dbName = os.Args[1]
	}
	// Connection established
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	log.Printf("Successfully connected to mongodb server at %v", url)

	//Connecting to the given database
	db := session.DB(dbName)
	if db == nil {
		log.Printf("db '%v' not found, exiting...", dbName)
		return
	}

	// iterate through collections
	cols, err := db.CollectionNames()
	if err != nil {
		log.Printf("No collections in db '%v'", dbName)
	}

	fmt.Printf("Collections in db '%v':\n", dbName)
	for i, v := range cols {
		fmt.Printf("[%v] - %v\n", i+1, v)
	}
}
