package main
import "github.com/TwiN/go-color"

import "fmt"

func HalfPyrm() {

	// *
	// **
	// ***
	// ****
	// *****
	// ******

	var el string = "* "

	for i := 0; i <= 6; i++ {
		// fmt.Println(el)
		for j := 0; j < i ; j++ {
			print(color.Ize(color.Blue, el))

		}
		fmt.Println()

	}

	fmt.Println()
	
}