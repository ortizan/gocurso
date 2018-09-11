package main

import "fmt"

type Persona struct {
	nombre string
	edad uint8
}

func main() {
	var p1 Persona 
	p1.nombre = "Antonio"
	p1.edad = 53
	fmt.Println(p1)

	p2 := Persona{}
	p2.nombre = "Irene"
	p2.edad = 18
	fmt.Println(p2)

	p3 := Persona{"Mari Paz", 55}
	fmt.Println(p3)
}
