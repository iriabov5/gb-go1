package main

import (
	"fmt"
)

func main() {
	array := []int{7, 4, 3, 2, 9, 12, 1, -2}
	insertionSort(array)
	fmt.Println(array)
}

func insertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		x := array[i]
		j := i
		for ; j >= 1 && array[j-1] > x; j-- {
			array[j] = array[j-1]
		}
		array[j] = x
	}
}
