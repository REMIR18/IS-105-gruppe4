package main

import "fmt"

func main() {
	const sample = "\xc2\xbd\x3f\x3d\x3f\xe2\x8c\x98"

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	//fmt.Printf("\n")

	//fmt.Println("Printf with %x:")
	//fmt.Printf("%x\n", sample)

	//fmt.Println("Printf with % x:")
	//fmt.Printf("% x\n", sample)

	//fmt.Println("Printf with %q:")
	//fmt.Printf("%q\n", sample)

	//fmt.Println("Printf with %+q:")
	//fmt.Printf("%+q\n", sample)

	fmt.Println("Printf with %s:")
	fmt.Printf("%s\n", sample)

	//fmt.Println("Printf with %c:")
	//fmt.Printf("%c\n", sample)
}
