package main

import "fmt"
import "time"

func main() {
	fmt.Println(time.Now().Weekday())
	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	fmt.Println("Calling x from main.")
	x()
	fmt.Println("Returned from x.")
}
func x() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in x", r)
		}
	}()
	fmt.Println("Executing x...")
	fmt.Println("Calling y.")
	y(0)
	fmt.Println("Returned normally from y.")
}

func y(i int) {
	fmt.Println("Executing y....")
	if i > 2 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in y", i)
	fmt.Println("Printing in y", i)
	y(i + 1)
}
