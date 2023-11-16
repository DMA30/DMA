package main

import "fmt"

func main() {
	empty := ""
	symbol := "#"
	space := " "
	a := 8
	b := 8
	for i := 0; i < a; i++ {
		for n := 0; n < b; n++ {
			if i%2 == 0 {
				if n%2 == 0 {
					empty = symbol
				} else {
					empty = space
				}
			} else {
				if n%2 == 0 {
					empty = space
				} else {
					empty = symbol
				}
			}
			fmt.Print(empty)
		}
		fmt.Println()
	}
}
