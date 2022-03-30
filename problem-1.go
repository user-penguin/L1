package main

import "fmt"

type Human struct {
	firstName string
	lastName  string
	thirdName string
	age       int
	height    int
	weight    int
}

func (h *Human) GetFullName() string {
	return h.lastName + " " + h.firstName + " " + h.thirdName
}

type Action struct {
	name string
	Human
}

func main() {
	action := Action{
		name: "test-1",
		Human: Human{
			firstName: "Dmitry",
			lastName:  "Kobzev",
			thirdName: "Igorevich",
			age:       25,
			height:    188,
			weight:    85,
		},
	}
	fmt.Printf("Full name of action's human: %s", action.GetFullName())
}
