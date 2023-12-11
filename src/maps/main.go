package main

import "fmt"

func main() {
	studentsAge := map[string]int{
		"john": 32,
		"bob":  31,
	}
	fmt.Println(studentsAge)

	// map without initialization
	studentsAge2 := make(map[string]int)
	fmt.Println(studentsAge2)

	// append elements
	studentsAge2["john"] = 32
	studentsAge2["bob"] = 31
	fmt.Println(studentsAge2)

	// printing an existing value
	fmt.Println("Bob's age is", studentsAge["bob"])

	// printing a non existing value, a default value is printed, no error thrown
	fmt.Println("Boby's age is", studentsAge["boby"])

	// checking if the value exists on the map or we are getting a default value
	age, exist := studentsAge["boby"]
	if exist {
		fmt.Println("Boby's age is", age)
	} else {
		fmt.Println("Boby's age couldn't be found")
	}

	// deleting an existing value
	delete(studentsAge, "john")
	fmt.Println(studentsAge)
	studentsAge["jhon"] = 32

	// deleting a non existing value, no error thrown
	delete(studentsAge, "boby")
	fmt.Println(studentsAge)

	// iterating over the map
	for name, age := range studentsAge {
		fmt.Printf("%s\t%d\n", name, age)
	}

	// iterating ignoring one of the values, oberve the "_" character instead of the name variable
	for _, age := range studentsAge {
		fmt.Printf("Ages %d\n", age)
	}

	// another way of getting only the value
	for name := range studentsAge {
		fmt.Printf("Names %s\n", name)
	}

}
