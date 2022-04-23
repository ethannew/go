package main

import (
	"fmt"
)

func main() {
	myArray := [5]string{"I", "am", "stupid", "and", "weak"}

	myArray[2] = "smart"
	myArray[4] = "strong"

	for _, value := range myArray {
		fmt.Println(value)
	}
}
