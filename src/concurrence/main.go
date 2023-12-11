package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkAPI(api string) {
	_, err := http.Get(api)
	if err != nil {
		fmt.Printf("ERROR: %s is down!\n", api)
		return
	}

	fmt.Printf("SUCCESS: %s is up and running!\n", api)
}

func checkAPIWithChannel(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api) // sends data throught the channel
		return
	}

	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api) // sends data throught the channel
}

func send(ch chan string, message string) {
	ch <- message
}

// separating the send and read functions

// here the channel is declared as a writting channel
func send2(ch chan<- string, message string) {
	fmt.Printf("Sending: %#v\n", message)
	ch <- message
}

// here the channel is declared as reading channel
func read(ch <-chan string) {
	fmt.Printf("Receiving: %#v\n", <-ch)
	//ch <- "Bye mother fucker" // this causes a compilation error, because the channel is just for reading not for sending
}

func main() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	for _, api := range apis {
		go checkAPI(api)
	}

	// has to sleep in order to wait for the prints from the go routine checkApi
	time.Sleep(3 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
	start = time.Now()
	// creating a channel in order to get data from the go routine
	ch := make(chan string)

	// in order to receive or send data use the next lines
	//ch <- x // sends (or writes ) x through channel ch
	//x = <-ch // x receives (or reads) data sent to the channel ch
	//<-ch // receives data, but the result is discarded (not saved on any variable, you can still print it)
	for _, api := range apis {
		go checkAPIWithChannel(api, ch)

	}

	for i := 0; i < len(apis); i++ {
		/*
					   here the program gets blocked until the data is received throught the channel,
					   then the program moves forward, so with the loop we can print all the go routine checkApi outputs.
			           be aware that this could block the programm if no data is received.
		*/
		fmt.Println(<-ch)
	}

	close(ch) // closes the channel

	elapsedChannel := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsedChannel.Seconds())

	/*  buffered stored channels, the data is stored as on a queue with the defined size,
	    so the program don't get blocked if the size is enough for all the calls 4 in this case.
	    With a size lower than 4 in this case the program would be blocked
	*/
	size := 4
	bufferedCh := make(chan string, size)
	// here the calls are not go routine, not recommended
	send(bufferedCh, "one")
	send(bufferedCh, "two")
	send(bufferedCh, "three")
	send(bufferedCh, "four")
	fmt.Println("All data sent to the buffered channel ...")

	for i := 0; i < size; i++ {
		fmt.Println(<-bufferedCh)
	}

	close(bufferedCh)

	fmt.Println("Done buffered channel!")

	// defining a lower size than the number of pieces of data we expect to receive, works fine because the last ones are go routines
	// with all the four calls without being go routines would block the programm because there will be only 2 slots
	size = 2
	bufferedCh2 := make(chan string, size)
	send(bufferedCh2, "one")
	send(bufferedCh2, "two")
	// these two calls are go routines, the recommended way when work with channels, the two ones above are not recommended
	go send(bufferedCh2, "three")
	go send(bufferedCh2, "four")
	fmt.Println("All data sent to the buffered channel 2 ...")

	for i := 0; i < 4; i++ {
		fmt.Println(<-bufferedCh2)
	}

	fmt.Println("Done buffered channel 2!")

	close(bufferedCh2)

	// defining a higher size than the number of pieces of data we expect to receive, works fine
	size = 6
	bufferedCh3 := make(chan string, size)
	go send(bufferedCh3, "one")
	go send(bufferedCh3, "two")
	// these two calls are go routines, the recommended way when work with channels, the two ones above are not recommended
	go send(bufferedCh3, "three")
	go send(bufferedCh3, "four")
	fmt.Println("All data sent to the buffered channel 3 ...")

	for i := 0; i < 4; i++ {
		fmt.Println(<-bufferedCh3)
	}

	fmt.Println("Done buffered channel 3!")

	// calling the separated send and read functions
	readingSendingChannel := make(chan string, 1)
	send2(readingSendingChannel, "Hello World!")
	read(readingSendingChannel)

}
