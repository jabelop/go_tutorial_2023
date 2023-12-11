package main

import (
	"fmt"
	"reflect"
)

func generateMultidimensionMatrix() [3][5]int {
	var twoD [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("Row", i, twoD[i])
	}
	return twoD
}

func printSliceWhileAppendingElements() {
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
	}
}

func main() {
	cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities:", cities, "length:", len(cities))

	// initialize an array without specifying the elments number
	cities2 := [...]string{"Wasingthon", "Ouckland", "Toluse", "Milan"}
	fmt.Println("Cities:", cities2, "length:", len(cities2))
	cities3 := append(cities2[0:], "Canberra")

	fmt.Println("Cities:", cities3, "typeof:", reflect.TypeOf(cities3))

	// initialize an array of 99 elements giving the -1 value to the last one
	numbers := [...]int{99: -1}
	fmt.Println("First Position:", numbers[0])
	fmt.Println("Last Position:", numbers[99])
	fmt.Println("Length:", len(numbers))

	fmt.Println("All at once", generateMultidimensionMatrix())

	// slices declaration and operations
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Println(months)
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))

	// Extracting sub slices
	quarter1 := months[0:3]
	quarter2 := months[3:6]
	quarter3 := months[6:9]
	quarter4 := months[9:12]

	// the cap value is the difference betwen the longitude of the original slice and the index chosen as initial
	// ex: for the quarter1 from months[0:3], months len 12, minus initial index 0, equal to 12.
	//     for the quarter2 from months[3:6], months len 12, minus initial index 3, equal to 9.
	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter3, len(quarter3), cap(quarter3))
	fmt.Println(quarter4, len(quarter4), cap(quarter4))

	// extending a slice with more capacity than len
	quarter2Extended := quarter2[:4]
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter2Extended, len(quarter2Extended), cap(quarter2Extended))

	printSliceWhileAppendingElements()

	// remove elements from a slice
	letters := []string{"A", "B", "C", "D", "E"}
	remove := 2

	if remove < len(letters) {

		fmt.Println("Before letters", letters, "Remove ", letters[remove])

		letters = append(letters[:remove], letters[remove+1:]...)

		fmt.Println("After letters", letters)
	}

	// manipulating slice without copy, all referenced slices are changed as well
	letters2 := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before2", letters2)

	slice1 := letters2[0:2]
	slice2 := letters2[1:4]

	slice1[1] = "Z"

	fmt.Println("After2", letters2)
	fmt.Println("Slice1", slice1)
	fmt.Println("Slice2", slice2)

	// manipulating slice with copy, all referenced slices are not changed
	letters3 := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before3", letters3)

	slice3 := letters3[0:2]

	slice4 := make([]string, 3)
	copy(slice4, letters[1:4])

	slice3[1] = "Z"

	fmt.Println("After3", letters3)
	fmt.Println("Slice3", slice3)
	fmt.Println("Slice4", slice4)
}
