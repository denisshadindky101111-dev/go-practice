package main

import "fmt"

func main() {
	// Задача 3:
	// Создать переменные:
	// b типа byte, smallI типа int32 и bigI типа uint64.
	// Присвоить каждой максимальное значение для ее типа,
	// затем добавить 1 и вывести результаты.
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b++
	smallI++
	bigI++

	fmt.Println("b =", b)
	fmt.Println("smallI =", smallI)
	fmt.Println("bigI =", bigI)
}
