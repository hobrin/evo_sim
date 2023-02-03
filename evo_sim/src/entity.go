package main

type Entity struct {
	x, y, r, rSq float64
}

func (e *Entity) updateRadius(r float64) {
	e.r = r
	e.rSq = r * r
}
