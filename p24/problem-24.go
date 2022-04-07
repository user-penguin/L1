/*Package p24
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/
package p24

import (
	"fmt"
	"math"
)

type point struct {
	x int
	y int
}

func NewPoint(xPos int, yPos int) *point {
	return &point{
		x: xPos,
		y: yPos,
	}
}

func (p *point) ComputeDistance(p2 point) float64 {
	xDelta, yDelta := p.x-p2.x, p.y-p2.y
	xSq, ySq := xDelta*xDelta, yDelta*yDelta
	return math.Sqrt(float64(xSq + ySq))
}

func Run() {
	p1, p2 := NewPoint(1, 2), NewPoint(20, -4)
	distance := p1.ComputeDistance(*p2)
	fmt.Printf("Расстояние между точками = %0.3f ед", distance)
}
