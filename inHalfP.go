package main

import "fmt"

func InvertedHalfPyr() {
	//  * * * * * *
	//  * * * * *
	//  * * * *
	//  * * *
	//  * *
	//  *

	var el string = "* "

	for i := 6; i >= 0; i-- {
		for j := 0; j < i ; j++ {
			fmt.Print(el)

		}
		fmt.Println()

	}
}
