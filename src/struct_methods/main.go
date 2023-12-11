package main

import (
	"fmt"
	"strings"
)

type triangle struct {
	size int
}

type square struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

func (t *triangle) doubleSize() {
	t.size *= 2
}

func (s square) perimeter() int {
	return s.size * 4
}

// defining methods for other types (string in this case)
type upperstring string

func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}

// mehod insertion (triangle type is inserted in a new struct along with its methods)
type coloredTriangle struct {
	triangle
	color string
}

/*
The method from the triangle inserted is automatically converted on a method of the coloredTriangle (the container struct) like this:
func (t coloredTriangle) perimeter() int {
    return t.triangle.perimeter()
}
So when coloredTriangle.perimeter() method is call,
this method is executed and his body calls the contained type (triangle) method, like a wraper.
*/

// overriding the triangle.perimeter() method so the new implementation is called with coloredTriangle.perimeter()
// if want to call the implementation from the inner triangle type you must do this: coloredTriangle.triangle.permiter()
type coloredTriangle2 struct {
	triangle
	color string
}

func (t coloredTriangle2) perimeter() int {
	return t.size * 3 * 2
}

func main() {
	t := triangle{3}
	s := square{4}
	fmt.Println("Perimeter (triangle):", t.perimeter())
	fmt.Println("Size (triangle):", t.size)
	fmt.Println("Perimeter (square):", s.perimeter())
	fmt.Println("Size (square):", s.size)

	t.doubleSize()
	fmt.Println("Size after doble size method call (triangle):", t.size)

	// method apply on other types (string in this case)
	str := upperstring("Learning Go!")
	fmt.Println(str)
	fmt.Println(str.Upper())

	// defining a new struct variable with another type inside (triangle in this case)
	trg := coloredTriangle{triangle{3}, "blue"}
	fmt.Println("Size:", trg.size)
	fmt.Println("Perimeter:", trg.perimeter())
	fmt.Println("Color:", trg.color)

	// overriding methods
	trg2 := coloredTriangle2{triangle{3}, "blue"}
	fmt.Println("Size:", trg2.size)
	fmt.Println("Perimeter (colored-container)", trg2.perimeter())
	fmt.Println("Perimeter (normal-inner)", trg2.triangle.perimeter())
}
