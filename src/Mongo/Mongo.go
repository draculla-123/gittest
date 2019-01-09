package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Structure in which it's gonna store the data
type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string        `json:"username"`
	Password string        `json:"password"`
	Email    string        `json:"email"`
}

func GetMongoSession() *mgo.Session {

	mgoSession, err := mgo.Dial("localhost")
	mgoSession.SetMode(mgo.Monotonic, true)

	if err != nil {
		log.Fatal("Failed to start the Mongo session.")
	}

	return mgoSession.Clone()
}

func CreateUser(u *User) (*User, error) {

	session := GetMongoSession()

	c := session.DB("poll").C("user")
	defer session.Close()

	err := c.Insert(u)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func main() {

	var user = User{
		Email:    "ayushm38@gmail.com",
		Password: "secret_i'm not telling ",
		Username: "draculla-123",
	}

	_, e := CreateUser(&user)
	if e != nil {
		panic(e)
	}

	log.Println("Done...")
}
