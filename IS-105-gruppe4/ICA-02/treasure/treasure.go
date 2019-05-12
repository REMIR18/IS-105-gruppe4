package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//henter inn filen
	t, err := ioutil.ReadFile("treasure.txt")
	if err != nil {
		fmt.Print(err)
	}

	//Gjør om inholdet i filen til string
	str := string(t)

	//printer stringen
	fmt.Printf(str)

	//Skriver ut stringen som ascii tekst
	fmt.Println("\n \n", "\x48\x65\x6c\x6c\x6f\x20\x65\x76\x65\x72\x79\x62\x6f\x64\x79\x20\x6f\x75\x74\x20\x74\x68\x65\x72\x65\x20\x75\x73\x69\x6e\x67\x20\x6d\x69\x6e\x69\x78\x20\x2d\x49\x27\x6d\x20\x64\x6f\x69\x6e\x67\x20\x61\x20\x28\x66\x72\x65\x65\x29\x20\x6f\x70\x65\x72\x61\x74\x69\x6e\x67\x20\x73\x79\x73\x74\x65\x6d\x20\x28\x6a\x75\x73\x74\x20\x61\x20\x68\x6f\x62\x62\x79\x2c\x20\x77\x6f\x6e\x27\x74\x20\x62\x65\x20\x62\x69\x67\x20\x61\x6e\x64\x70\x72\x6f\x66\x65\x73\x73\x69\x6f\x6e\x61\x6c\x20\x6c\x69\x6b\x65\x20\x67\x6e\x75\x29\x20\x66\x6f\x72\x20\x33\x38\x36\x28\x34\x38\x36\x29\x20\x41\x54\x20\x63\x6c\x6f\x6e\x65\x73\x2e\x20\x20\x54\x68\x69\x73\x20\x68\x61\x73\x20\x62\x65\x65\x6e\x20\x62\x72\x65\x77\x69\x6e\x67\x73\x69\x6e\x63\x65\x20\x61\x70\x72\x69\x6c\x2c\x20\x61\x6e\x64\x20\x69\x73\x20\x73\x74\x61\x72\x74\x69\x6e\x67\x20\x74\x6f\x20\x67\x65\x74\x20\x72\x65\x61\x64\x79\x2e\x20\x20\x49\x27\x64\x20\x6c\x69\x6b\x65\x20\x61\x6e\x79\x20\x66\x65\x65\x64\x62\x61\x63\x6b\x20\x6f\x6e\x74\x68\x69\x6e\x67\x73\x20\x70\x65\x6f\x70\x6c\x65\x20\x6c\x69\x6b\x65\x2f\x64\x69\x73\x6c\x69\x6b\x65\x20\x69\x6e\x20\x6d\x69\x6e\x69\x78\x2c\x20\x61\x73\x20\x6d\x79\x20\x4f\x53\x20\x72\x65\x73\x65\x6d\x62\x6c\x65\x73\x20\x69\x74\x20\x73\x6f\x6d\x65\x77\x68\x61\x74\x73\x61\x6d\x65\x20\x70\x68\x79\x73\x69\x63\x61\x6c\x20\x6c\x61\x79\x6f\x75\x74\x20\x6f\x66\x20\x74\x68\x65\x20\x66\x69\x6c\x65\x2d\x73\x79\x73\x74\x65\x6d\x20\x28\x64\x75\x65\x20\x74\x6f\x20\x70\x72\x61\x63\x74\x69\x63\x61\x6c\x20\x72\x65\x61\x73\x6f\x6e\x73\x29\x61\x6d\x6f\x6e\x67\x20\x6f\x74\x68\x65\x72\x20\x74\x68\x69\x6e\x67\x73\x29\x2e\x20\x49\x27\x76\x65\x20\x63\x75\x72\x72\x65\x6e\x74\x6c\x79\x20\x70\x6f\x72\x74\x65\x64\x20\x62\x61\x73\x68\x28\x31\x2e\x30\x38\x29\x20\x61\x6e\x64\x20\x67\x63\x63\x28\x31\x2e\x34\x30\x29\x2c\x20\x61\x6e\x64\x20\x74\x68\x69\x6e\x67\x73\x20\x73\x65\x65\x6d\x20\x74\x6f\x20\x77\x6f\x72\x6b\x2e\x54\x68\x69\x73\x20\x69\x6d\x70\x6c\x69\x65\x73\x20\x74\x68\x61\x74\x20\x49\x27\x6c\x6c\x20\x67\x65\x74\x20\x73\x6f\x6d\x65\x74\x68\x69\x6e\x67\x20\x70\x72\x61\x63\x74\x69\x63\x61\x6c\x20\x77\x69\x74\x68\x69\x6e\x20\x61\x20\x66\x65\x77\x20\x6d\x6f\x6e\x74\x68\x73\x2c\x20\x61\x6e\x64\x49\x27\x64\x20\x6c\x69\x6b\x65\x20\x74\x6f\x20\x6b\x6e\x6f\x77\x20\x77\x68\x61\x74\x20\x66\x65\x61\x74\x75\x72\x65\x73\x20\x6d\x6f\x73\x74\x20\x70\x65\x6f\x70\x6c\x65\x20\x77\x6f\x75\x6c\x64\x20\x77\x61\x6e\x74\x2e\x20\x20\x41\x6e\x79\x20\x73\x75\x67\x67\x65\x73\x74\x69\x6f\x6e\x73\x61\x72\x65\x20\xf8\x77\x65\x6c\x63\x6f\x6d\x65\x2c\x20\x62\x75\x74\x20\x49\x20\x77\x6f\x6e\x27\x74\x20\x70\x72\x6f\x6d\x69\x73\x65\x20\x49\x27\x6c\x6c\x20\xe6\x69\x6d\x70\x6c\x65\x6d\x65\x6e\x74\x20\x74\x68\x65\x6d\x20\x3a\x2d\x29\xe5")
}
