package main

import "fmt"

type human struct {
	firstName string
	lastName  string
	thirdName string
	age       int
	height    int
	weight    int
}

func (h *human) GetFullName() string {
	return h.lastName + " " + h.firstName + " " + h.thirdName
}

type action struct {
	human
	name string
}

func main() {
	action := action{
		name: "test-1",
		human: human{
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
