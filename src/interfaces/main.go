package main

import (
	"fmt"
	"math"
)

// defining an interface
type Shape interface {
	Perimeter() float64
	Area() float64
}

// implementing an interface
type Square struct {
	size float64
}

// defining the interface methods for a type lique Square means it is implementing the interface
func (s Square) Area() float64 {
	return s.size * s.size
}

func (s Square) Perimeter() float64 {
	return s.size * 4
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printInformation(s Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
	fmt.Println()
}

// Stringer interface implementation
type Stringer interface {
	String() string
}

type Person struct {
	Name, Country string
}

func (p Person) String() string {
	return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}

func main() {
	var s Shape = Square{3}
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
	fmt.Println()

	// calling a method which expects an Shape interface as argument
	printInformation(s)

	c := Circle{6}
	printInformation(c)

	// use of the Stringer interface, the Printf method will use our String() implementation
	rs := Person{"John Doe", "USA"}
	ab := Person{"Mark Collins", "United Kingdom"}
	fmt.Printf("%s\n%s\n", rs, ab)

}
