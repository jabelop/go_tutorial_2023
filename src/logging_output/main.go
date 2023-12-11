package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// regular log
	log.Print("Hey, I'm a log!")

	// log fatal, logs and stops the program execution limke with os.exit(1)
	//#log.Fatal("Hey, I'm an error log!")
	fmt.Println("Can you see me?")

	// the same as above with panic adding call stack info, comment the fatal line in order to see this output
	//log.Panic("Hey, I'm an error log!")
	fmt.Println("Can you see me?")

	// adding a prefix to logs
	log.SetPrefix("main(): ")
	log.Print("Hey, I'm a log!")
	//#log.Fatal("Hey, I'm an error log!")

	// logging into a file
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("Hey, I'm a log!")

	// look for https://github.com/sirupsen/logrus or https://github.com/rs/zerolog for more advances features
	// for installation:  go get -u github.com/rs/zerolog/log
}
