package main

import (
	"fmt"
)

func quickSort(arr []uint32) []uint32 {
	newArr := make([]uint32, len(arr))

	for i, v := range arr {
		newArr[i] = v
	}

	sort(newArr, 0, len(arr)-1)

	return newArr
}

func sort(arr []uint32, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if arr[i] < pivot {
			temp := arr[splitIndex]

			arr[splitIndex] = arr[i]
			arr[i] = temp

			splitIndex++
		}
	}

	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	sort(arr, start, splitIndex-1)
	sort(arr, splitIndex+1, end)
}

func main() {
	list := []uint32{41, 63, 64, 72, 78, 85, 75, 38, 82, 3, 99, 50, 83, 37, 94, 8, 87, 1, 40, 100, 15, 6, 58, 51, 45, 43, 5, 73, 53, 4, 74, 11, 60, 18, 36, 67, 76, 9, 84, 19, 57, 30, 90, 80, 55, 32, 29, 52, 35, 54}

	s := quickSort(list)

	fmt.Println(s)
	fmt.Println(list)
}
