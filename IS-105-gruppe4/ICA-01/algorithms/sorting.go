package algorithms

import (
	"fmt"
)

var toBeSorted [5]int = [5]int{8, 5, 3, 9, 2}

// Les https://en.wikipedia.org/wiki/Bubble_sort
func Bubble_sort_modified(input [5]int) {
	// Deres kode her

}

// Implementering av Bubble_sort algoritmen
func Bubble_sort(toBeSorted [5]int) {
	// find the length of list n
	n := len(toBeSorted)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if toBeSorted[j] > toBeSorted[j+1] {
				temp := toBeSorted[j+1]
				toBeSorted[j+1] = toBeSorted[j]
				toBeSorted[j] = temp
			}
		}
	}
	fmt.Println(toBeSorted)
}
func main() {
	fmt.Println("These numbers are sorted")
	Bubble_sort(toBeSorted)
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
