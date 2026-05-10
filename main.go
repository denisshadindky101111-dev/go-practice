package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func saveTodoList(tasks []string) error {
	return os.WriteFile("todo.txt", []byte(strings.Join(tasks, "\n")), 0644)
}

func loadTodoList() ([]string, error) {
	f, err := os.Open("todo.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var list []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		if line != "" {
			list = append(list, line)
		}
	}
	return list, s.Err()
}

func main() {
	fmt.Println("Welcome to the Todo List App")
	fmt.Println("1. Add a task")
	fmt.Println("2. List tasks")
	fmt.Println("3. Remove a task")
	fmt.Println("4. Exit")

	todoList, err := loadTodoList()
	if err != nil && !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "warning: could not load todo.txt: %v\n", err)
		todoList = nil
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter a command:")
		scanner.Scan()
		command := scanner.Text()

		switch command {
		case "1":
			fmt.Println("Enter a task:")
			scanner.Scan()
			task := scanner.Text()
			todoList = append(todoList, task)
			if err := saveTodoList(todoList); err != nil {
				fmt.Fprintf(os.Stderr, "could not save: %v\n", err)
			} else {
				fmt.Println("Task added successfully")
			}
		case "2":
			fmt.Println("Todo List:")
			for i, task := range todoList {
				fmt.Printf("%d. %s\n", i+1, task)
			}
		case "3":
			fmt.Println("Enter the task number to remove:")
			scanner.Scan()
			taskNumber := scanner.Text()
			taskNumberInt, err := strconv.Atoi(taskNumber)
			if err != nil {
				fmt.Println("Invalid task number")
				continue
			}
			if taskNumberInt > 0 && taskNumberInt <= len(todoList) {
				todoList = append(todoList[:taskNumberInt-1], todoList[taskNumberInt:]...)
				if err := saveTodoList(todoList); err != nil {
					fmt.Fprintf(os.Stderr, "could not save: %v\n", err)
				} else {
					fmt.Println("Task removed successfully")
				}
			} else {
				fmt.Println("Invalid task number")
			}
		case "4":
			fmt.Println("Exiting...")
			if err := saveTodoList(todoList); err != nil {
				fmt.Fprintf(os.Stderr, "could not save: %v\n", err)
			}
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}
