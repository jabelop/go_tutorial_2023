package main

import (
	"fmt"

	"methodvisibility/geometry"
)

func main() {
	t := geometry.Triangle{}
	t.SetSize(3)
	fmt.Println("Perimeter", t.Perimeter())
}
