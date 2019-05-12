package main

import (
	"fmt"

	fileutils "github.com/IS-105-gruppe4/IS-105-gruppe4/ICA-03/fileutils/fileutilsgo"
)

// Henter tekstfil og bruker FileToByteslice funksjonen for så å printe ut filene
func main() {
	file1 := fileutils.FileToByteslice("files/text1.txt")
	file2 := fileutils.FileToByteslice("files/text2.txt")

	fmt.Println(file1)
	fmt.Println(file2)
}
