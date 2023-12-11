package main

import "fmt"

func main() {
	val := 0
	for true {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)
		if val < 0 {
			panic("Number less than 0")
		} else if val == 0 {
			fmt.Println("0 is less than 1")
		} else {
			fmt.Println("You entered:", val)
		}

	}
}
