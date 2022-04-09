/*Package p21
Реализовать паттерн «адаптер» на любом примере.
*/
package p21

//WheelMachine интерфейс, по которому должна работать наша система
type WheelMachine interface {
	TurnWheelsRight()
	TurnWheelsLeft()
}

func Run() {
	car := Car{}
	tank := Tank{}
	car.TurnWheelsRight()
	car.TurnWheelsLeft()
	adapter := NewAdapter(&tank)
	adapter.TurnWheelsRight()
	adapter.TurnWheelsLeft()
}
