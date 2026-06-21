package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Приветствие ===")

	var name string
	var age int

	fmt.Print("Введите ваше имя: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println("Ошибка ввода имени:", err)
		return
	}

	fmt.Print("Введите ваш возраст: ")
	_, err = fmt.Scan(&age)
	if err != nil {
		fmt.Println("Ошибка ввода возраста:", err)
		return
	}

	fmt.Printf("Привет, %s! Вам %d лет.\n", name, age)

	fmt.Println("=== До свидания ===")

	fmt.Scanln()
	fmt.Println("\nНажмите Enter для выхода...")
	fmt.Scanln()
}
