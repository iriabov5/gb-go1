package main

import (
	"fmt"
	"time"
)

var mapFib = map[int]int{}

func main() {
	var n int
	fmt.Println("Введите число: ")
	fmt.Scanln(&n)
	start1 := time.Now()
	fmt.Printf("%d число Фибоначчи равно :%d\n", n, simpleFib(n))
	fmt.Println("Время вычисления первым способом: ", time.Since(start1))
	start2 := time.Now()
	fmt.Printf("%d число Фибоначчи равно :%d\n", n, mapRecurFib(n))
	fmt.Println("Время вычисления вторым способом: ", time.Since(start2))

}

func simpleFib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return simpleFib(n-1) + simpleFib(n-2)
}

func mapRecurFib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	a, aExists := mapFib[n]
	if !aExists {
		mapFib[n] = mapRecurFib(n-1) + mapRecurFib(n-2)
		return mapFib[n]
	} else {
		return a
	}
}
