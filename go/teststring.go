package main

import (
	"fmt"
)

type A struct {
	Name string
	Age  int
}

func (a A) String() string {
	return fmt.Sprintf("the name is %s, the age is %d", a.Name, a.Age)
}

func main() {
	s := "this is a prefix string "
	a := A{
		Name: "Bob",
		Age:  15,
	}
	fmt.Println(a)
	ss := s + a.String()
	fmt.Println(ss)
}
