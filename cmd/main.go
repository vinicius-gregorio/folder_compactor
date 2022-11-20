package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	os.Setenv("GITHUB_TOKEN", "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")
	flag.Usage = func() {
		fmt.Printf("Usage: %s \n\n \t\t go run filesLocation newFilesLocation. \n\n Example: \t go run ./cmd C:/Users/USER/Downloads/compactor_test C:/Users/USER/Downloads/compactor_test/compacted", os.Args[0])
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
