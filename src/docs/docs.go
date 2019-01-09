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
	//Scanner function basically it sees the length of the command line argument
	if 1 == len(os.Args) {
		log.Printf("No db specified, using '%v'", dbName)
	} else {
		dbName = os.Args[1]
	}
	//Connection established
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	log.Printf("Successfully connected to mongodb server at %v", url)

	// Connected to the given database
	db := session.DB(dbName)
	if db == nil {
		log.Printf("db '%v' not found, exiting...", dbName)
		return
	}

	// iterate through collections
	fmt.Printf("Collections in db '%v':\n", dbName)
	cols, err := db.CollectionNames()
	if err != nil {
		return
	}

	for _, c := range cols {
		fmt.Printf("[%v]\n", c)
		listDocs(db, c)
	}
}

//iterating through documents
func listDocs(db *mgo.Database, col string) {
	coll := db.C(col)
	if coll == nil {
		return
	}

	//we use interface when we dont know the exact structure of the result to be displayed
	var result []interface{}
	coll.Find(nil).All(&result)
	for i, d := range result {
		fmt.Printf("\tDoc%2v - %v\n", i+1, d)
	}
}
