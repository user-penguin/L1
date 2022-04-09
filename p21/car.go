package p21

import "fmt"

//Car обычный автомобиль, реализующий методы интерфейса колёсной техники
type Car struct {
	WheelMachine
}

func (c *Car) TurnWheelsRight() {
	fmt.Println("Машина повернула колёса и поехала вправо")
}

func (c *Car) TurnWheelsLeft() {
	fmt.Println("Машина повернула колёса и поехала влево")
}
