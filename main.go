package main

import (
	"fmt"
)

func main() {
	// Instancia 1 del objeto "Course"
	Go := Course{
		Name:    "Go desde Cero",
		Price:   12.34,
		IsFree:  false,
		UserIDs: []uint{12, 56, 89},
		Classes: map[uint]string{
			1: "Introduccion",
			2: "Structures",
			3: "Functions",
		},
	}
	fmt.Println(Go.Name) // Se puede traer uno o varios valores de la instancia

	// Instancia 2 del Objeto "Course"
	CSS := Course{
		Name:   "CSS Desde Cero",
		IsFree: true,
	}
	fmt.Printf("%v\n", CSS)

	// Instancia 3 del Objeto "Course"

	React := Course{}

	React.Name = "React JS From Scratch"
	React.IsFree = false
	React.Price = 14.60
	React.Classes = map[uint]string{
		1: "NPM",
		2: "Modules",
	}

	Go.ChangePrice(90.17)
	fmt.Println(Go.Price)

}
