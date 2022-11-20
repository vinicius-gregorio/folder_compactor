package main

import (
	"errors"
	"log"
	"os"
)

func CompactFilesToFolder(location string, newLocation string) error {
	// CreateDirectoryIfNotExists(newLocation)
	if _, err := os.Stat(newLocation); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(newLocation, os.ModePerm)
		if err != nil {
			log.Fatal(err)

			return err
		}
	}
	f, err := os.Open(location)
	if err != nil {
		log.Fatal(err)

		return err
	}
	files, err := f.Readdir(0)
	if err != nil {
		log.Fatal(err)

		return err
	}

	for _, v := range files {
		if !v.IsDir() {
			err := os.Rename(location+"/"+v.Name(), newLocation+"/"+v.Name())
			if err != nil {
				log.Fatal(err)
				return err
			}
		}
	}
	return nil
}
