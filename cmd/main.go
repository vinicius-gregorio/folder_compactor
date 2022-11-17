package main

import (
	"fmt"
	"os"
)

func main() {
	// find where my downloads folder is and create a os.open

	f, err := os.Open("C:/Users/Gregorio/Downloads/compactor_test")
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range files {
		fmt.Println(v.Name(), v.IsDir())
	}
}
