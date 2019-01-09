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

	//entering the name of the database u wanna work with
	fmt.Println("Enter the name of the database you wanna work with and if it is not their it will be created")
	var s string
	fmt.Scanf("%s", &s)

	//checking for valid database
	db := session.DB(s)
	if db == nil {
		log.Printf("db '%v' not found, exiting...", s)
		return
	}

	// iterate through collections
	cols, err := db.CollectionNames()
	if err != nil {
		log.Printf("No collections in db '%v'", s)
	}

	fmt.Printf("Collections in db '%v':\n", s)
	for i, v := range cols {
		fmt.Printf("[%v] - %v\n", i+1, v)
	}

	//entering the name of the collection u wanna work with
	fmt.Println("Enter the name of the collection you wanna work with and if it is not their it will be created")
	var c string
	fmt.Scanf("%s", &c)

	//iterating through documents
	coll := db.C(c)
	if coll == nil {
		return
	}

	var result []interface{}
	coll.Find(nil).All(&result)
	for i, d := range result {
		fmt.Printf("\tDoc%2v - %v\n", i+1, d)
	}

}
