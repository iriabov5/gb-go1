package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	fmt.Print("Введите число a: ")
	fmt.Scanln(&a)

	var b int
	fmt.Print("Введите число b: ")
	fmt.Scanln(&b)

	var op string
	fmt.Print("Введите операцию: +, -, *, / , %, ! , **: ")
	fmt.Scanln(&op)
	fmt.Print("Результат: ")
	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("деление на ноль")
		} else {
			a := float64(a)
			b := float64(b)
			fmt.Printf("%.2f", a/b)
		}
	case "%":
		fmt.Println(a % b)
	case "!":
		fmt.Printf("факториал числа %d равен %d", a, factorial(uint64(a)))
	case "**":
		fmt.Printf("%d в степени %d равно %d", a, b, int(math.Pow(float64(a), float64(b))))
	default:
		fmt.Println("ошибочная операция")
	}
}

func factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}
