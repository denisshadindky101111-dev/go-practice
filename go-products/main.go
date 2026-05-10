package main

import (
	"fmt"
)

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

func (p Product) TotalPrice() float64 {
	return p.Price * float64(p.Quantity)
}

func (p Product) ProductInfo() {
	fmt.Printf("ID: %d, Name: %s, Price: %f, Quantity: %d, Total Price: %f\n", p.ID, p.Name, p.Price, p.Quantity, p.TotalPrice())
}

func AddProduct(sliceProducts []Product, p Product) []Product {
	sliceProducts = append(sliceProducts, p)
	return sliceProducts
}

func GetProductById(sliceProducts []Product, id int) int {
	for i, p := range sliceProducts {
		if p.ID == id {
			return i
		}
	}
	return -1
}

func UpdateQuantity(sliceProducts *[]Product) {
	var id, newQty int
	fmt.Print("Введите ID товара: ")
	fmt.Scan(&id)
	fmt.Print("Введите новое количество: ")
	fmt.Scan(&newQty)

	idx := GetProductById(*sliceProducts, id)
	if idx == -1 {
		fmt.Println("Товар не найден.")
		return
	}
	(*sliceProducts)[idx].Quantity = newQty
	fmt.Printf("Количество обновлено: товар ID %d, quantity = %d\n", id, newQty)
}

func PrintAllProducts(sliceProducts *[]Product) {
	for _, p := range *sliceProducts {
		p.ProductInfo()
	}
}	

func main() {

	var sliceProducts []Product
	currentId := 1

	for {
		fmt.Println("1. Добавить товар")
		fmt.Println("2. Обновить количество")
		fmt.Println("3. Вывести информацию о всех товарах")
		fmt.Println("4. Выйти")
		var choice int
		fmt.Scan(&choice)

	switch choice {
	case 1:
		var name string
		var price float64
		var quantity int
		fmt.Print("Введите название товара: ")
		fmt.Scan(&name)
		fmt.Print("Введите цену товара: ")
		fmt.Scan(&price)
		fmt.Print("Введите количество товара: ")
		fmt.Scan(&quantity)
		newProduct := Product{ID: currentId, Name: name, Price: price, Quantity: quantity}
		currentId++
		sliceProducts = AddProduct(sliceProducts, newProduct)
	case 2:
		UpdateQuantity(&sliceProducts)
	case 3:
		PrintAllProducts(&sliceProducts)
	case 4:
		fmt.Println("Выход из программы...")
		return
	}
}
}
