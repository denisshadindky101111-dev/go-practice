package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Knetic/govaluate"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== Консольный калькулятор выражений ===")
	fmt.Println("Введите математическое выражение (например: (2 + 2) * 3)")
	fmt.Println("Для выхода введите: exit")
	fmt.Println("----------------------------------------")

	for {
		fmt.Print("Введите выражение: ")
		if !scanner.Scan() {
			break
		}

		expression := scanner.Text()
		if expression == "exit" {
			break
		}
	
		if expression == "" {
			continue
		}

		expr, err := govaluate.NewEvaluableExpression(expression)
		if err != nil {
			fmt.Printf("Ошибка синтаксиса: %v\n", err)
			continue
		}

		result, err := expr.Evaluate(nil)
		if err != nil {
			fmt.Printf("Ошибка вычисления: %v\n", err)
			continue
		}

		fmt.Printf("Результат: %v\n", result)
	}

}
