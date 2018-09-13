package main

import "fmt"

/*func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}	*/

func main() {
	/* for i := 5; i > 0; i-- {
		fmt.Println(i)
	}
	*/

	for i := 4; i >= 0; i-- {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}
 