package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <csvFile>\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	userInput, err := getUserInputs()
	if err != nil {
		log.Fatal(err)
	}
	err = CompactFilesToFolder(userInput.Location, userInput.NewLocation)
	if err != nil {
		log.Fatal(err)
	}
}
