package main

import (
	"fmt"

	"github.com/tuvace/is105-ica02/unicode/unicode"
)

func main() {

	input := "\xE2\x80\x9C\x6E\x6F\x72\x64\x20\x6F\x67\x20\x73\xC3\xB8\x72\xE2\x80\x9D\xE2\x80\x8B"
	outputIsland := unicode.Translate(input, "is")
	outputJpan := unicode.Translate(input, "jp")

	fmt.Println(outputIsland)
	fmt.Println(outputJpan)

	//Test av UnicodeCodePointDemo
	unicode.UnicodeCodePointDemo()

}
