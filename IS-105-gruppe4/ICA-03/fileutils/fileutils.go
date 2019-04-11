package fileutils

import (
	"io"
	"log"
	"os"
)

func FileToByteslice("text1.txt" string) []byte {

	// Open file for reading
	file, err := os.Open("text1.txt")

	if err != nil {
		log.Fatal(err)
	}
	finfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size_of_slice := finfo.Size()

	// The file.Read() function can read a
	// tiny file into a large byte slice,
	// but io.ReadFull() will return an
	// error if the file is smaller than
	// the byte slice
	byteSlice := make([]byte, size_of_slice)

	_, err = io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	return byteSlice

}
