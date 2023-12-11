package main

import (
	"fmt"
	"time"
)

var quit = make(chan bool)

func fib(c chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Done calculating Fibonacci!")
			return
		}
	}
}

func main() {
	start := time.Now()

	command := ""
	data := make(chan int)

	// doing this way instead of calling the function on the loop controls the order of the elements,
	// look at the fib body to see how it keeps the loop contained making use of the channels

	go fib(data)

	for {

		num := <-data
		fmt.Println(num)
		//fmt.Scanf("%s", &command)
		if num > 70 {
			quit <- true
			break
		}
		if command == "quit" {
			quit <- true
			break
		}
	}

	time.Sleep(1 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
