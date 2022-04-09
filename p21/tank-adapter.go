package p21

type Adapter struct {
	*Tank
}

func NewAdapter(tank *Tank) WheelMachine {
	return &Adapter{tank}
}

func (a *Adapter) TurnWheelsRight() {
	a.BlockRightTrack()
}

func (a *Adapter) TurnWheelsLeft() {
	a.BlockLeftTrack()
}
