package main

import (
	"fmt"
)

func partition(a []int, lo, hi int) int {
	p := a[hi]
	for j := lo; j < hi; j++ {
		if a[j] < p {
			a[j], a[lo] = a[lo], a[j]
			lo++
		}
	}

	a[lo], a[hi] = a[hi], a[lo]
	return lo
}

func quickSort(a []int, lo, hi int) {
	if lo > hi {
		return
	}

	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func main() {
	list := []int{41, 63, 64, 72, 78, 85, 75, 38, 82, 3, 99, 50, 83, 37, 94, 8, 87, 1, 40, 100, 15, 6, 58, 51, 45, 43, 5, 73, 53, 4, 74, 11, 60, 18, 36, 67, 76, 9, 84, 19, 57, 30, 90, 80, 55, 32, 29, 52, 35, 54}

	quickSort(list, 0, len(list)-1)
	fmt.Println(list)
}
