package main

import (
	"bufio"
	"fmt"
	"github/Quaqmre/MarsRover/command"
	"github/Quaqmre/MarsRover/compass"
	"github/Quaqmre/MarsRover/rover"
	"log"
	"os"
	"strconv"
	"strings"

	"github/Quaqmre/MarsRover/plato"
)

func main() {
	fmt.Println("Hello \nWelcome to my Code Challenge for HepsiBurada \nMars Rover \nMy Name is Mehmet Akif Tatar")

	// Plato,rover ve yönlendirme bilgileri alınır.
	platoX, platoY, roverX, roverY, roverDirct, commands := ReadStdIn()

	//Alınan bilgiler ile command oluşturulur.
	command := command.NewCommand(commands)

	//Alınan bilgiler ile pusula oluşturulur
	compass, err := compass.NewCompass(roverDirct)
	if err != nil {
		log.Fatal(err)
	}

	//Alınan bilgiler ile Rover oluşturulur.
	rover, err := rover.NewRover(roverX, roverY, roverDirct)

	//Rovera pusula ve emir bilgileri yüklenir.
	rover.SetCommands(command)
	rover.SetCompass(compass)

	if err != nil {
		log.Fatal(err)

	}

	//Aracın üzerinde hareket edeceği plato oluşturulur
	plato := plato.NewPlato(platoX, platoY, rover)

	//Plato üzerinde arayış başlar
	x, y, d := plato.Search()

	//Binary yön tekrar 'N','W'.. gibi anlaşılır yönlere çevrilir
	direct := BinaryToDirection(d)

	fmt.Println("Solve:")
	fmt.Println(x, y, direct)

	fmt.Println("Good job, Rover")

}

func BinaryToDirection(i int) string {
	direction := ""
	switch i {
	case 1:
		direction = "W"
	case 1 << 1:
		direction = "S"
	case 1 << 2:
		direction = "E"
	case 1 << 3:
		direction = "N"

	}
	return direction
}

// Console üzerinden okumamız için kullanılan function
func ReadStdIn() (int, int, int, int, rune, string) {
	scanner := bufio.NewScanner(os.Stdin)
	stdin1 := ""
	stdin2 := ""
	stdin3 := ""

	for i := 1; i <= 3; i++ {
		if i == 1 {
			fmt.Println("Get plato coordinates like,5 5")
			scanner.Scan()
			stdin1 = scanner.Text()
		} else if i == 2 {
			fmt.Println("Get rover coordinates like,1 2 N")
			scanner.Scan()
			stdin2 = scanner.Text()
		} else {
			fmt.Println("Get Commands like,LMLMLMLMM")
			scanner.Scan()
			stdin3 = scanner.Text()
		}
	}
	//---------------Input2------------------------

	arStr := strings.Split(stdin1, " ")
	if len(arStr) != 2 {
		log.Fatal("first input should be like 5 5")
	}
	px, err := strconv.Atoi(arStr[0])

	if err != nil {
		log.Fatal(err)
	}

	py, err2 := strconv.Atoi(arStr[1])

	if err2 != nil {
		log.Fatal(err2)
	}

	//---------------Input2------------------------
	arStr1 := strings.Split(stdin2, " ")
	if len(arStr1) != 3 {
		log.Fatal("first input should be like 1 2 N")
	}

	rx, err := strconv.Atoi(arStr1[0])
	if err != nil {
		log.Fatal(err)
	}
	ry, err := strconv.Atoi(arStr1[1])
	if err != nil {
		log.Fatal(err)
	}
	direction := rune(arStr1[2][0])

	//---------------Input3------------------------
	commands := stdin3

	return px, py, rx, ry, direction, commands
}
