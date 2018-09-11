package main

import "fmt"

func main() {

	/* nombres := make([]string, 0)
	nombres = append(nombres, "Antonio")
	nombres = append(nombres, "Mari Paz")
	fmt.Println(nombres)
	fmt.Printf("La longitud de nombres es: %d y su capacidad es %d\n", len(nombres), cap(nombres))
	nombres = append(nombres, "Irene")
	*/
	nombres := []string {
		"Antonio",
		"Mari Paz",
		"Maripacilla",
		"Irene",
		"javi",
	}
	fmt.Println(nombres)
	fmt.Printf("La longitud de nombres es: %d y su capacidad es %d\n", len(nombres), cap(nombres))


}