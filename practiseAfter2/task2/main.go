package main

import "fmt"

const value = 20

func main() {
	// Задача 2:
	// Объявить константу value, которая может быть присвоена
	// как целочисленной переменной, так и переменной с плавающей запятой.
	// Присвоить ее i и f, затем вывести значения.
	var i int = value
	var f float64 = value

	fmt.Println("i =", i)
	fmt.Println("f =", f)
}
