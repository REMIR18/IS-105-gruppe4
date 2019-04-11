package main

import (
    "fmt"
    "log"
    "io/ioutil"
    "unicode/utf8"
    "strconv"
)

func main() {
    file, err := os.Open("file.txt")
    if err != nil {
        log.Fatal(err)

        
// Initializing my scanner
type Scanner interface {
    countChar(text string) int

    frequentSym(text string) // Return value is not yet implemented

    Scan()

    Run()
}

/* method counting characters */
func countChar(sc Scanner, text string) int { ... }

func frequentSym(sc Scanner, text string) {
    // Make a map with string key and integer value
    symbols := make(map[string] int)

    // Iterate through each char in text
    for _, sym := range text {
        // Convert rune to string
        char := strconv.QuoteRune(sym)
        // Set this string as a key in map and assign a counter value 
        count := symbols[char]

        if count == symbols[char] {
            // increment the value
            symbols[char] = count + 1
        } else {
            symbols[char] = 1
        }
    }
}
}