package main

import "log"

const (
	LOCATION     = "C:/Users/Gregorio/Downloads/compactor_test"
	NEW_LOCATION = "C:/Users/Gregorio/Downloads/compactor_test/compacted"
)

func main() {

	err := CompactFilesToFolder(LOCATION, NEW_LOCATION)
	if err != nil {
		log.Fatal(err)
	}
}
