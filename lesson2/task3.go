package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите трехзначное число: ")
	fmt.Scanln(&a)
	fmt.Println("Количество сотен равно: ", a/100)
	fmt.Println("Количество десятков равно: ", a%100/10)
	fmt.Println("Количество единиц равно: ", a%100%10)
}
