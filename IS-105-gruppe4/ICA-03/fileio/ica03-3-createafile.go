package main

import (
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
)

func main() {
	//create a new file
	newFile, err = os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()

	//truncate a file
	err := os.Truncate("test.txt", 100)
	if err != nil {
		log.Fatal(err)
	}

}
