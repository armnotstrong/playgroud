package main

import "fmt"

type xx struct {
	Name string
	Age  int
}

func main() {
	var people *xx
	people = &xx{
		Name: "Bob",
		Age:  15,
	}
	fmt.Println(*people)
	var x string
	fmt.Println(x)
}
