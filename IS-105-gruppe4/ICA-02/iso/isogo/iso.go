package iso

import "fmt"

//Extended ascii tabell
const asciiExtendedLiteral = "\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8a\x8b\x8c\x8d\x8e\x8f" +
	"\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9a\x9b\x9c\x9d\x9e\x9f" +
	"\xa0\xa1\xa2\xa3\xa4\xa5\xa6\xa7\xa8\xa9\xaa\xab\xac\xad\xae\xaf" +
	"\xb0\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe\xbf" +
	"\xc0\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf" +
	"\xd0\xd1\xd2\xd3\xd4\xd5\xd6\xd7\xd8\xd9\xda\xdb\xdc\xdd\xde\xdf" +
	"\xe0\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xeb\xec\xde\xee\xef" +
	"\xf0\xf1\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xfb\xfc\xfd\xfe\xff"

// Oppgave 2a
// Lag selv en "string literal" med utvidet ASCII
// Blir det kun en slik "string literal" eller trenger man flere
// for å representere utvidet ASCII?

// IterateOverExtendedASCIIStringLiteral tar en "string literal" som INN-data
func IterateOverExtendedASCIIStringLiteral() {
	// Kode for Oppgave 2a
	for i := 128; i < 255; i++ {
		fmt.Printf("%X %q %b\n", i, i, i)
	}
}

// IterateOverExtendedASCIIStringLiteralTXT skriver ut tegn til stringen txt
func IterateOverExtendedASCIIStringLiteralTXT(txt string) {
	for i := 0; i < len(txt); i++ {
		fmt.Printf("%X, %q, %b\n", txt[i], txt[i], txt[i])
	}
}

// GetExtendedASCIIStringLiteral henter konstanten extended asciiExtendedLiteral
func GetExtendedASCIIStringLiteral() string {
	return asciiExtendedLiteral
}

// GreetingExtendedASCII returnerer en tekst-streng i
// utvidet ASCII
// Kode for Oppgave 2c
func GreetingExtendedASCII() {
	fmt.Printf("Salut, ça va °-) Ça coute €50")
	fmt.Printf("Salut, ça va °-) Κοστίζει €50")
}

//PrintASCII prints ascii
func PrintASCII() {
	for i := 128; i <= 255; i++ {
		fmt.Printf("%X, %q, %b\n", i, i, i)
	}
}
