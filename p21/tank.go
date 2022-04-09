package p21

import "fmt"

//Tank инородное тело с гусиницами, которое как-то надо повернуть по законам колёсной техники (Adaptee)
type Tank struct {
}

func (t *Tank) BlockRightTrack() {
	fmt.Println("Танк начинал поворачивать вправо")
}

func (t *Tank) BlockLeftTrack() {
	fmt.Println("Танк начинал поворачивать Влево")
}
