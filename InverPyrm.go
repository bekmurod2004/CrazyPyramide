package main

import "fmt"

func INvPyrm() {
	// unpiramide

		var el string = " * "
	var el2 string = "  *"

	for i := 6; i > -1; i-- {
		for j := 0; j < 6; j++ {
			if i%2 == 0 {
				if i == 0 {
					if j ==2 {
						fmt.Print(el2)
					}else {
						fmt.Print("   ")
					}

				}

				if i == 2 {
					if j == 1 ||j == 2 ||j == 3 {
						fmt.Print(el2)
					}else {
						fmt.Print("   ")
					}
				}

				if i == 4 {
					if j == 0 || j ==1 || j ==2|| j ==3|| j ==4  {
						fmt.Print(el2)
					}else {
						fmt.Print("   ")
					}
				}

			} else {
				if i == 1 {
					if j ==2 || j ==3 {
	                    fmt.Print(el)
	                }else {
						fmt.Print("   ")
					}

				}

				if i == 3 {
					if j ==1 || j ==2|| j ==3|| j ==4 {
	                    fmt.Print(el)
	                }else {
						fmt.Print("   ")
					}

				}

				if i == 5 {
					fmt.Print(el)
				}
			}
		}
		fmt.Println()
	}

	fmt.Println()
	
}