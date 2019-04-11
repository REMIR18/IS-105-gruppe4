package main

import (
	"bufio"
	"io"
	"os"
	"testing"
)

func BenchmarkReadFileBuffered(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f, err := os.Open(". pg100.txt")
		if err != nil {
			panic(err)
		}

		r := bufio.NewReader(f)

		_, err = r.ReadString('\n')
		for err == nil {
			_, err = r.ReadString('\n')
		}
		if err != io.EOF {
			panic(err)
		}

		f.Close()
	}
}
