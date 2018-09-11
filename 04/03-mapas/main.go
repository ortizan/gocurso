package main

import "fmt"

func main() {
	/*
	idiomas := make(map[string]string)
	idiomas["es"] = "Español"
	idiomas["en"] = "Inglés"
	idiomas["fr"] = "Frances"
	*/
	idiomas := map[string]string {"es":"Español", "en":"Inglés", "fr":"Frances"}
	fmt.Println(idiomas)
	fmt.Println(idiomas["es"])
	delete(idiomas, "fr")
	fmt.Println(idiomas)
}