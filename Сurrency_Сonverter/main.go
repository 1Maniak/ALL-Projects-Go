package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	// "golang.org/x/text/transform"
	"golang.org/x/text/encoding/charmap"
)

// СТРУКТУРА ДЛЯ XML
type Valute struct {
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

// ЧИСТКА КОНСОЛИ
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

// ПРИВЕТСТВЕННОЕ МЕНЮ
func showMenu() {
	fmt.Println("\n╔════════════════════════════╗")
	fmt.Println("║   КОНВЕРТЕР ВАЛЮТ v0.1.1   ║")
	fmt.Println("║         by Maniak          ║")
	fmt.Println("╚════════════════════════════╝")
	fmt.Println()
	fmt.Println("========================")
	fmt.Println("1. Конвертировать")
	fmt.Println("2. Выход")
	fmt.Println("========================")
	fmt.Print("Выберите опцию: ")
}

func showCurrencies() {
	fmt.Println("\n╔════════════════════════════╗")
	fmt.Println("║      ДОСТУПНЫЕ ВАЛЮТЫ:     ║")
	fmt.Println("╚════════════════════════════╝")
	fmt.Println("\n1. Доллар США (USD)")
	fmt.Println("2. Евро (EUR)")
	fmt.Println("3. Российский рубль (RUB)")
	fmt.Println("==============================")
}

func showLoading() {
	fmt.Print("Загрузка курсов валют")
	for i := 0; i < 4; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println()
}

func getCurrencyRate(currencyCode string) (float64, error) {
	// ЗАПРОС К САЙТУ ЦБ
	url := "https://www.cbr-xml-daily.ru/daily.xml"

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return 0, fmt.Errorf("не удалось подключиться: %v", err)
	}
	defer resp.Body.Close()

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("неподдерживаемая кодировка: %s", charset)
	}

	// РАЗБИРАЕМ XML
	var valCurs ValCurs

	err = decoder.Decode(&valCurs)
	if err != nil {
		return 0, fmt.Errorf("ошибка при разборе XML: %v", err)
	}

	// ИЩЕМ ВАЛЮТУ
	for _, valute := range valCurs.Valutes {
		if valute.CharCode == currencyCode {
			// ЗАМЕНА ЗАПЯТОЙ НА ТОЧКУ
			valueStr := strings.Replace(valute.Value, ",", ".", 1)
			rate, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				return 0, fmt.Errorf("ошибка при преобразовании курса: %v", err)
			}
			return rate, nil
		}
	}
	return 0, fmt.Errorf("валюта %s не найдена", currencyCode)
}

func convertCurrency(amount float64, from int, to int) (float64, error) {
	fromCode := getCurrencyCode(from)
	toCode := getCurrencyCode(to)

	// fromRate, err := getCurrencyRate(fromCode)
	// if err != nil {
	// 	return 0, err
	// }

	// toRate, err := getCurrencyRate(toCode)
	// if err != nil {
	// 	return 0, err
	// }

	fromRate := 1.0
	toRate := 1.0

	var err error

	if fromCode != "RUB" {
		fromRate, err = getCurrencyRate(fromCode)
		if err != nil {
			return 0, err
		}
	}

	if toCode != "RUB" {
		toRate, err = getCurrencyRate(toCode)
		if err != nil {
			return 0, err
		}
	}

	inRub := amount * fromRate
	result := inRub / toRate

	return result, nil
}

func getCurrencyCode(currencyNum int) string {
	switch currencyNum {
	case 1:
		return "USD"
	case 2:
		return "EUR"
	case 3:
		return "RUB"
	default:
		return ""
	}
}

// ОСНОВНАЯ ФУНКЦИЯ
func main() {
	for {
		clearScreen()
		showMenu()

		var option int
		fmt.Scan(&option)

		if option == 2 {
			break
		}

		if option == 1 {
			clearScreen()
			showCurrencies()

			var amount float64
			var from, to int

			fmt.Print("Введите сумму: ")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Ошибка! Сумма должна быть больше 0!")
				fmt.Println("Нажмите Enter для продолжения...")
				fmt.Scanln()
				fmt.Scanln()
				continue
			}

			fmt.Print("Из какой валюты (1-3): ")
			fmt.Scan(&from)
			if from < 1 || from > 3 {
				fmt.Println("Ошибка! Выберите валюту от 1 до 3!")
				fmt.Println("\nНажмите Enter для продолжения...")
				fmt.Scanln()
				fmt.Scanln()
				continue
			}

			fmt.Print("В какую валюту (1-3): ")
			fmt.Scan(&to)
			if to < 1 || to > 3 {
				fmt.Println("Ошибка! Выберите валюту от 1 до 3!")
				fmt.Println("\nНажмите Enter для продолжения...")
				fmt.Scanln()
				fmt.Scanln()
				continue
			}

			if from == to {
				fmt.Println("Ошибка! Нельзя выбрать одиннаковые валюты!")
				fmt.Println("\nНажмите Enter для продолжения...")
				fmt.Scanln()
				fmt.Scanln()
				continue
			}

			showLoading()

			result, err := convertCurrency(amount, from, to)
			if err != nil {
				fmt.Println("\nОшибка при получениии курса: ", err)
				fmt.Println("Проверьте подключение к интернету!")
				fmt.Println("\nНажмите Enter для продолжения...")
				fmt.Scanln()
				fmt.Scanln()
				continue
			}

			fmt.Println("Готово!")

			fromName := ""
			toName := ""

			switch from {
			case 1:
				fromName = "USD"
			case 2:
				fromName = "EUR"
			case 3:
				fromName = "RUB"
			}

			switch to {
			case 1:
				toName = "USD"
			case 2:
				toName = "EUR"
			case 3:
				toName = "RUB"
			}

			fmt.Println("\n==============================")
			fmt.Printf("%.2f %s = %.2f %s\n", amount, fromName, result, toName)
			fmt.Println("==============================")

			fmt.Println("\nНажмите Enter для продолжения...")
			fmt.Scanln()
			fmt.Scanln()

		}
	}

}
