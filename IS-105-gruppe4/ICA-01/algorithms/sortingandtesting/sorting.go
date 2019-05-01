package algorithms

import "fmt"

//Kode lånt fra anzuh på guthub.com
// Les https://en.wikipedia.org/wiki/Bubble_sort
func Bubble_sort_modified(list []int) []int {
	// set n to length of array to sort
	n := len(list)

	// first for pass goes from 0 to length of array
	for i := 0; i < n; i++ {
		// second for pass goes from 0 to length of array -1
		for j := 0; j < n-1; j++ {
			// check if value at j index is smaller than value at j + 2
			if list[j] < list[j+1] {
				// set temp var to j+1
				temp := list[j+1]
				// set value at index j+1 to previous value
				list[j+1] = list[j]
				// set original value to temp
				list[j] = temp
				// effectively, swap values j and j+1
			}
			fmt.Println(list)
		}
	}
	// return array so it can be printed or set to a var or whatever lmao
	return list
}

// Implementering av Bubble_sort algoritmen
func Bubble_sort(list []int) {
	// find the length of list n
	n := len(list)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j+1]
				list[j+1] = list[j]
				list[j] = temp
			}
		}
	}
}

// Implementering av Quicksort algoritmen
func QSort(values []int) {
	qsort(values, 0, len(values)-1)
}

func qsort(values []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := values[l]
	i := l + 1

	for j := l; j <= r; j++ {
		if pivot > values[j] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[l], values[i-1] = values[i-1], pivot

	qsort(values, l, i-2)
	qsort(values, i, r)
}
