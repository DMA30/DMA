package main

import "fmt"

func main() {
	var a int
	var b int
	empty := ""
	symbol := "#"
	space := " "

	fmt.Print("Укажите длинну поля:")
	fmt.Scan(&b)
	fmt.Print("Укажите высоту поля:")
	fmt.Scan(&a)

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
