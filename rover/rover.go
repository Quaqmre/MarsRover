package rover

import (
	"errors"
	"github/Quaqmre/MarsRover/command"
	"github/Quaqmre/MarsRover/compass"

	"log"
)

// Marsta görevli Aracımızdır.
type Rover struct {
	Compass  *compass.Compass
	cordX    int
	cordY    int
	Commands *command.Command
}

//Belli parametreler ile Rover oluşturmamızı sağlar
func NewRover(x, y int, d rune) (rover *Rover, err error) {
	rover = &Rover{}
	rover.Compass, err = compass.NewCompass(d)
	if err != nil {
		return nil, err
	}
	rover.cordX = x
	rover.cordY = y
	return rover, nil
}

//Rover üzerine yerleştirilen yön bulma algoritmasıdır.
func (r *Rover) SetCompass(c *compass.Compass) {
	if c == nil {
		log.Fatal("SetCompass - can not be null")
	}
	r.Compass = c
}

//Gelen emirler Rover'imize yüklenir.
func (r *Rover) SetCommands(cm *command.Command) {
	if cm == nil {
		log.Fatal("SetCompass - can not be null")
	}
	r.Commands = cm
}

// Gelen emirleri 2 katagoride değerlendirmemizi sağlar
func (r *Rover) ResolveCommand(c rune) error {
	switch c {
	case 'L', 'R':
		r.Compass.TurnCompass(c)
		return nil
	case 'M':
		err := r.Move()
		if err != nil {
			return err
		}
		return nil
	default:
		log.Println("Wrong Command: ", string(c))
		return errors.New("ResolveCommand - wrong Command")
	}
}

// Roverin bulunduğu durumun bilgisini döner.
func (r *Rover) GetCordinant() (int, int, int) {
	return r.cordX, r.cordY, r.Compass.Direction
}

// Plato üzerinde arama yapmak hareket algoritmasını oluşturur.
func (r *Rover) Move() error {

	switch r.Compass.Direction {
	case 1:
		r.cordX -= 1
	case 1 << 2:
		r.cordX += 1
	case 1 << 3:
		r.cordY += 1
	case 1 << 1:
		r.cordY -= 1
	default:
		log.Println("I can't Move :)")
		return errors.New("out of range")

	}
	return nil
}
