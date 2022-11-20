package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type UserInput struct {
	Location    string
	NewLocation string
}

func getUserInputs() (UserInput, error) {
	if len(os.Args) < 2 {
		return UserInput{}, errors.New("a filepath argument is required")
	}

	flag.Parse()

	fileLocation := flag.Arg(0)
	fileNewLocation := flag.Arg(1)

	fmt.Println("File location:", fileLocation)
	fmt.Println("File location:", flag.Arg(1))

	return UserInput{fileLocation, fileNewLocation}, nil
}
