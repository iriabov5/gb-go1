package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64
	fmt.Println("Введите площадь круга: ")
	fmt.Scanln(&s)
	fmt.Println("Диаметр круга равен: ", 2*math.Sqrt(s/math.Pi))
	fmt.Println("Длина окружности равна: ", math.Sqrt(4*s*math.Pi))
}
