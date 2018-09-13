package main

import "fmt"

/* func main() {
	numeros := []string{"cero","uno", "dos", "tres"}
	for i, v := range numeros {
		fmt.Println(i, v)
	}
}*/

/*func main() {
	numeros := []string {"cero", "uno", "dos", "tres"}
	for _ , v := range numeros{
		fmt.Println(v)
	}

}
*/

/*func main() {
	paises := map[string]string{"es": "España", "fr":"Francia", "uk":"Reino Unido"}
	for i, v := range paises {
		fmt.Println(i, v)
	}
}
*/

func main() {
	letras := "España"
	for _, v := range letras {
		fmt.Println(string(v))
	}
}