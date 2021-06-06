package plato

import (
	"github/Quaqmre/MarsRover/rover"
)

// Marsta bulunan platoyu temsil eder
type Plato struct {
	lowerX int
	lowerY int
	highX  int
	highY  int

	r *rover.Rover
}

func NewPlato(maxX, maxY int, r *rover.Rover) *Plato {
	plato := Plato{
		lowerX: 0,
		lowerY: 0,
		highX:  maxX,
		highY:  maxY,
		r:      r,
	}
	return &plato
}

// Plato üzerinde gezecek olan roveri bulur.
func (p *Plato) SetRover(r *rover.Rover) {
	p.r = r
}

// Plato üzerinde bulunan Roverin arama görebini başlatır.
func (p *Plato) Search() (int, int, int) {

	for p.r.Commands.GetNextCommand() == nil {
		r := p.r.Commands.GetCurrentCommand()
		if r == 'M' {
			p.r.Move()
			continue
		}
		p.r.Compass.TurnCompass(r)
	}

	return p.r.GetCordinant()
}
