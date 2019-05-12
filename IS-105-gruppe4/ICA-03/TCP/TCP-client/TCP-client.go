package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	i := 0 //connections to server
	for {
		c, err := net.Dial("tcp", "127.0.0.1:5000") // net.Dial brukes for å prøve å koble til porten 5000
		if err != nil {
			fmt.Println(err, "Could not connect to the server.")
			os.Exit(1)
		} else {
			i++
			if i <= 1 {
				fmt.Println("Connected to server")
			} else {
				fmt.Println("Reconnected to server")
			}
		}

	}

	//lager en ny bufio reader som heter connbuff, den skal lese av stringen

	for {
		connbuf := bufio.NewReader(c)
		str, err := connbuf.ReadString('\n') //leser av stringen
		if len(str) > 0 {                    //hvis lengden på stringen er større enn 0, vil den bli printet ut
			fmt.Println(c, str+"\n") // printer ut
		}
		if err != nil { //hvis feil ikke er lik null, vil koblingsforsøket avsluttes
			break
		}

	}
}
