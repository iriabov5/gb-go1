package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("Введите длину стороны a: ")
	fmt.Scanln(&a)
	fmt.Println("Введите длину стороны b: ")
	fmt.Scanln(&b)
	fmt.Println("Площадь прямоугольника равна: ", a*b)
}
