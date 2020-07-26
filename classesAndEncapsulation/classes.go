package classesAndEncapsulation

import "math"

type Point2d struct {
	x float64
	y float64
}

type Point interface {
	distanceToOrigin() float64
}

func (p Point2d) NewPoint() Point2d {
	return Point2d{
		x: 0,
		y: 0,
	}
}

func (p Point2d) distanceToOrigin() float64 {
	t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
	return math.Sqrt(t)
}

func main() {
	println("Classes")
	println("-------")
	p1 := Point2d{
		x: 5,
		y: 5,
	}
	println(p1.distanceToOrigin())
}
