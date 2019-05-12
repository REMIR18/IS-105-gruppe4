package lineshift

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//Bruker kode fra oppgave 1a på starten
func LineShiftUsed(diktnavn string) []byte {
	// Open file for reading
	file, err := os.Open(diktnavn)

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
	//sjekker om diktet bruker linefeed eller carriage return/linefeed
	d := string(byteSlice[:])
	fmt.Println("Dette diktet ble skrevet i Mac eller Windows:")
	fmt.Println("Denne filen er skrevet i Mac:", strings.Contains(d, "\n"))
	fmt.Println("Denne filen er skrevet i Windows:", strings.Contains(d, "\r\n"))

	return byteSlice

}
