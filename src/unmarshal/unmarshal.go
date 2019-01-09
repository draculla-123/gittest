package main

import "encoding/json"
import "fmt"

//Person Structure of all the people present
type Person struct {
	Name   string
	Fruits string
}
type Person1 struct {
	Name   string
	Fruits []string
}

type response1 struct {
	Page   int
	Fruits []string
}

func main() {

	m := Person{"alice", "grapes"}
	res1B, _ := json.Marshal(m)
	fmt.Println(string(res1B))

	res1D := &response1{1, []string{"apple", "peach", "pear"}}
	res1B1, _ := json.Marshal(res1D)
	fmt.Println(string(res1B1))
	fmt.Println(m, res1D)

	byt := []byte(`{"Name":"alex","Fruits":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	result := Person1{}
	json.Unmarshal(byt, &result)
	fmt.Println(result)

}
