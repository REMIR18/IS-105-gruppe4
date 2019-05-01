package slice

//import "fmt"

// AllocateVar har INN-argument b
// b - antall bytes brukeren ønsker å allokere
// Returnerer en slice av type []byte
//
func AllocateVar(b int) []byte { //funksjonen deklarer en variabel som et heltall, og typen byte brukes til å tildele en minneposisjon til et element.
	var x []byte
	x = make([]byte, b) //setter lengde og kapasitet blir satt til b
	return x
}

// AllocateMake tar lengde og kapasitet som b og lager en ny slice
//
func AllocateMake(b int) []byte { //funksjonen lager en ny slice
	slice := make([]byte, b)
	return slice //returnerer den nyopprettede slicen
}

// Reslice takes a slice and reslices it
func Reslice(slc []byte, lidx int, uidx int) []byte { //
	reslice := AllocateMake(len(slc)) //bruker AllocateMake til å opprette en ny slice med samme lengde som slc
	reslice = slc[lidx:uidx]          // reslice bytter ut den orginale slicen
	slc = reslice
	return slc //returnerer den nye slicen
}
