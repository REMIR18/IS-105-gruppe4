package iso

import (
	"fmt"
	"testing"
)

//Tester om konstant teksten presenteres i ASCII extended.
func TestGreetingExtendedASCII(t *testing.T) {
	IterateOverExtendedASCIIStringLiteralTXT(GetExtendedASCIIStringLiteral())
}

//Tekst som blir testet.
const Test = "Salut, ça va °-) Ça coute €50" +
	"Salut, ça va °-) Κοστίζει €50" + "Salut, ça va °-) Κοστίζει €50 Forstår du?"

//Tester om tallet ikke er extended ASCII og printer ut hvis tallet ikke er en del av extended ASCII.
func TestASCII(t *testing.T) {
	for i := 0; i < len(Test); i++ {
		if Test[i] >= 255 {
			t.Fail()
			fmt.Println("Loop number;", i, "This number is not a part of extended ascii", Test[i])
		}
		if Test[i] <= 128 {
			t.Fail()
			fmt.Println("Loop number;", i, "This number is not a part of extended ascii", Test[i])
			fmt.Println(Test[i])
		}
	}
}
