package main

import "log"

const (
	LOCATION     = "C:/Users/Gregorio/Downloads/compactor_test"
	NEW_LOCATION = "C:/Users/Gregorio/Downloads/compactor_test/compacted"
)

func main() {

	// f, err := os.Open("C:/Users/Gregorio/Downloads/compactor_test")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// files, err := f.Readdir(0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// if _, err := os.Stat(NEW_LOCATION); errors.Is(err, os.ErrNotExist) {
	// 	err := os.Mkdir(NEW_LOCATION, os.ModePerm)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// }

	// for _, v := range files {
	// 	fmt.Println(v.Name(), v.IsDir())
	// 	if !v.IsDir() {
	// 		err := os.Rename(LOCATION+"/"+v.Name(), NEW_LOCATION+"/"+v.Name())
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 	}
	// }

	// err := CreateDirectoryIfNotExists(NEW_LOCATION)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	zz := CompactFilesToFolder(LOCATION, NEW_LOCATION)
	if zz != nil {
		log.Fatal(zz)
	}
}
