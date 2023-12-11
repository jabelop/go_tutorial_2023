package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

type Person struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

type Employee2 struct {
	Person
	ManagerID int
}

type Contractor struct {
	Person
	CompanyID int
}

type Person2 struct {
	ID        int
	FirstName string `json:"name"` // tag for specifying a different property name for the json
	LastName  string
	Address   string `json:"address,omitempty"` // tag for specifying a different property name and not to include empty values for this property
}

type Employee3 struct {
	Person2
	ManagerID int
}

type Contractor2 struct {
	Person2
	CompanyID int
}

func main() {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	fmt.Println(employee)
	employeeCopy := &employee
	employeeCopy.FirstName = "David"
	fmt.Println(employee)

	// structs insertion
	employee2 := Employee2{
		Person: Person{
			FirstName: "John",
		},
	}
	employee2.LastName = "Doe"
	fmt.Println(employee2.FirstName)

	// structs json encode decode with tags for field logic (look the struct)

	employees := []Employee3{
		Employee3{
			Person2: Person2{
				LastName: "Doe", FirstName: "John", Address: "Elm St.",
			},
		},
		Employee3{
			Person2: Person2{
				LastName: "Campbell", FirstName: "David",
			},
		},
	}

	data, _ := json.Marshal(employees)
	fmt.Printf("%s\n", data)

	var decoded []Employee3
	json.Unmarshal(data, &decoded)
	fmt.Printf("%v", decoded)
	fmt.Println()
}
