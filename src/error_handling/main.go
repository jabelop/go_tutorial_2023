package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrNotFound = errors.New("Employee not found!") // Custom error

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	employee, err := getInformation(1000)
	if err != nil {
		// Something is wrong. Do something.
		fmt.Println(err)
	} else {
		fmt.Println(employee)
	}

	employee2, err := getInformationJustReturningTheErrorIfThereWasOne(1000)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(employee2)
	}

	employee3, err := getInformationWithErrorContext(1000)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(employee3)
	}

	employee4, err := getInformationWithRetrievals(1000)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(employee4)
	}

}

func getInformation(id int) (*Employee, error) {
	employee, err := apiCallEmployee(id)
	return employee, err
}

func getInformationJustReturningTheErrorIfThereWasOne(id int) (*Employee, error) {
	employee, err := apiCallEmployee(id)
	if err != nil {
		return nil, err // Simply return the error to the caller.
	}
	return employee, nil
}

func getInformationWithErrorContext(id int) (*Employee, error) {
	employee, err := apiCallEmployee(id)
	if err != nil {
		return nil, fmt.Errorf("Got an error when getting the employee information: %v", err)
	}
	return employee, nil
}

func getInformationWithRetrievals(id int) (*Employee, error) {
	for tries := 0; tries < 3; tries++ {
		employee, err := apiCallEmployee(1000)
		if err == nil {
			return employee, nil
		}

		fmt.Println("Server is not responding, retrying ...")
		time.Sleep(time.Second * 2)
	}

	return nil, fmt.Errorf("server has failed to respond to get the employee information")
}

func apiCallEmployee(id int) (*Employee, error) {
	if id != 1001 {
		return nil, ErrNotFound
	}
	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}
