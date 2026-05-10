package main

import "fmt"

func main() {
	// Задача 1:
	// Объявить целочисленную переменную i со значением 20.
	// Присвоить значение i переменной с плавающей запятой f.
	// Вывести значения i и f.
	var i int = 20
	var f float64 = float64(i)

	fmt.Println("i =", i)
	fmt.Println("f =", f)
}
