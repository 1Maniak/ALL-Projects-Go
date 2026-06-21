package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func showMenu() {
	fmt.Println("╔════════════════════════════╗")
	fmt.Println("║   ГЕНЕРАТОР ПАРОЛЕЙ v1.0   ║")
	fmt.Println("║         by Maniak          ║")
	fmt.Println("╚════════════════════════════╝")
	fmt.Println()
	fmt.Println("========================")
	fmt.Println("1. Сгенерировать пароль")
	fmt.Println("2. Выход")
	fmt.Println("========================")
	fmt.Print("Выберите опцию: ")
}

func getPasswordLength() int {
	for {
		var length int
		fmt.Print("Введите длину пароля: ")
		_, err := fmt.Scanln(&length)
		if err != nil {
			fmt.Println("Ошибка, введите число.")
			fmt.Scanln()
			continue
		}

		if length > 0 {
			return length
		}
		fmt.Println("Длинна должна быть больше 0!")
	}
}

func getUserChoice() int {
	for {
		var choice int
		fmt.Println("========================")
		fmt.Println("Выберите тип пароля: ")
		fmt.Println("1 - только цифры")
		fmt.Println("2 - только буквы")
		fmt.Println("3 - буквы + цифры")
		fmt.Println("4 - буквы + цифры + символы")
		fmt.Println("========================")
		fmt.Print("Ваш выбор: ")

		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Ошибка: Введите число от 1 до 4!")
			fmt.Scanln()
		} else if choice < 1 {
			fmt.Println("Ошибка: Число не может быть меньше 1!")
		} else if choice > 4 {
			fmt.Println("Ошибка: Число не может быть больше 4!")
		} else {
			return choice
		}
	}
}

func generatePassword(length int, choice int) string {
	var charset string
	switch choice {
	case 1:
		charset = "0123456789"
	case 2:
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case 3:
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	case 4:
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
	default:
		charset = "0123456789"
	}

	rand.Seed(time.Now().UnixNano())
	password := ""

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		randomChar := charset[randomIndex]
		password = password + string(randomChar)
	}
	return password
}

// func showGeneratingEffect() {
// fmt.Print("Генерация пароля")
// for i := 0; i < 3; i ++ {
// time.Sleep(300 * time.Millisecond)
// fmt.Print(".")
// }
// fmt.Println()
// }

func showGeneratingEffect() {
	chars := []string{"|", "/", "-", "\\"}
	fmt.Print("Генерация пароля")
	for i := 0; i < 10; i++ {
		fmt.Printf("\rГенерация пароля %s", chars[i%len(chars)])
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println("\rГенерация пароля: Готово!")
}

func main() {
	fmt.Println("=== Generator Password ===")
	for {
		clearScreen()
		showMenu()
		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			length := getPasswordLength()
			choice := getUserChoice()
			showGeneratingEffect()

			password := generatePassword(length, choice)
			fmt.Println("========================================")
			fmt.Println("Сгенерированный пароль: ", password)
			fmt.Println("========================================")

			fmt.Scanln()
			fmt.Print("Нажмите Enter для выхода...")
			fmt.Scanln()

		case 2:
			fmt.Println("=== Generator Password ===")
			return

		default:
			fmt.Println("Неверный выбор! Поробуйте снова.")
			fmt.Println("Нажмите Enter для продолжения...")
			fmt.Scanln()
			fmt.Scanln()
		}
	}
}
