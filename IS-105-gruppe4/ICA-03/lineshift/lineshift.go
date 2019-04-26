package main

//import packages that the system need to work
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//argument that provides access to command-line arguemnts.
func main() {
	arg1 := os.Args[1]
	MacOrWindowsLineShift(arg1)

}

// MacOrWindowsLineShift opens a file, and checking the OS
func MacOrWindowsLineShift(filname string) {
	file1, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	//Reads a file
	data1, err := ioutil.ReadAll(file1)
	if err != nil {
		log.Fatal(err)
	}
	text1Platform := "Windows"
	if !match1 {
		text1Platform = "Mac"
	}

	fmt.Print("The textfile is using %s line breaks\n", text1Platform)

}
