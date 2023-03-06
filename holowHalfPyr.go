package main

import "fmt"

func HollowInvertedHalfPyr() {

	var el string = " *"
	var pro string = "  "

		// 666666
	// 55555
	// 4444
	// 333
	// 22
	// 1

	// 012345
	// 01234
	// 0123
	// 012
	// 01
	// 0

	// @ @ @ @ @ @
	// @       @
	// @     @
	// @   @
	// @ @
	// @
	for i := 6; i >= 0; i-- {
		for j := 0; j < i ; j++ {
			if i !=5 && i != 4 && i != 3{
				fmt.Print(el)

			}
			if i == 5 {
				if j == 1  || j == 2 || j == 3 {
					fmt.Print(pro)
				}else {
					fmt.Print(el)
				}

			}

			if i == 4 {
				if j == 1  || j == 2  {
					fmt.Print(pro)
				}else {
					fmt.Print(el)
				}

			}

			if i == 3 {
				if j == 1 {
					fmt.Print(pro)
				}else {
					fmt.Print(el)
				}

			}

		}
		fmt.Println()

	}
	
}