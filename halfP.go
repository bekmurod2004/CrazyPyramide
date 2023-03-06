package main

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
			fmt.Print(el)

		}
		fmt.Println()

	}

	fmt.Println()
	
}