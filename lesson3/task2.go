package main

import "fmt"

func main() {
	var n int
	fmt.Println("Введите число n: ")
	fmt.Scanln(&n)
	printPrimeNumbers(n)

}

func printPrimeNumbers(n int) {
	fmt.Printf("Простые числа от 1 до %d:\n", n)
	for i := 1; i <= n; i++ {
		if isNumberPrime(i) {
			fmt.Println(i)
		}
	}
}

func isNumberPrime(n int) bool {
	for i := 1; i < n; i++ {
		if i != 1 && n%i == 0 {
			return false
		}
	}
	return true
}
