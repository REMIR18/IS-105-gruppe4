package main

import (
	"fmt"

	fileutils "github.com/IS-105-gruppe4/IS-105-gruppe4/ICA-02/fileutils/filefolder"
)

func main() {
	// Bruker funkjsonen FileToByteslice på hver av de dte lang filene
	lang01 := fileutils.FileToByteslice("../files/lang01.wl")
	lang02 := fileutils.FileToByteslice("../files/lang02.wl")
	lang03 := fileutils.FileToByteslice("../files/lang03.wl")

	// Skriver ut hver av filene med %c
	// som skriver ut med det utvidet ASCII kodesettet
	fmt.Printf("%c", lang01)
	fmt.Printf("%c", lang02)
	fmt.Printf("%c", lang03)
}
