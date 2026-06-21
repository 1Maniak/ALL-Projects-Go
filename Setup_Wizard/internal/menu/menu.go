package menu

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"setup_wizard/internal/stages"

	"github.com/AlecAivazis/survey/v2"
)

// Глобальные переменные
var (
	WarningMessage    string
	BrokerInfoMessage string
)

// Очистка экрана
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

// Главное меню
type Menu struct {
	stages []stages.Stage
}

// Создает новое меню
func NewMenu(s []stages.Stage) *Menu {
	return &Menu{
		stages: s,
	}
}

// Показывает меню и обрабатывает выбор
func (m *Menu) Show() error {
	for {
		clearScreen()

		// Предупржедение
		if WarningMessage != "" {
			fmt.Println(WarningMessage)
		}

		// Информация о брокере
		if BrokerInfoMessage != "" {
			fmt.Println(BrokerInfoMessage)
		}
		fmt.Println()

		fmt.Println("========================================")
		fmt.Println("    	  Setup_Wizard v0.1.0			 ")
		fmt.Println("========================================")
		fmt.Println()

		// Список пунктов
		var options []string
		for _, stage := range m.stages {
			status := stage.Status()
			name := stage.Name()
			options = append(options, fmt.Sprintf("%s %s", status, name))
		}
		options = append(options, "Завершить настройку")

		// Показываем меню
		var choice string
		prompt := &survey.Select{
			Message: "Выберите действие:",
			Options: options,
		}

		err := survey.AskOne(prompt, &choice, survey.WithFilter(func(filter string, value string, index int) bool {
			return true
		}))
		if err != nil {
			return err
		}

		// Проверяем не выбран ли выход
		if choice == "Завершить настройку" {
			fmt.Println("До свидания!")
			return nil
		}

		// Находим выбранный этап
		for _, stage := range m.stages {
			status := stage.Status()
			name := stage.Name()
			if choice == fmt.Sprintf("%s %s", status, name) {
				log.Printf("Выбран этап: %s", name)
				if err := stage.Run(); err != nil {
					fmt.Printf("Ошибка выполнения этапа: %v\n", err)
					fmt.Println("Нажмите Enter для продолжения...")
					fmt.Scanln()
					continue
				}
			}
		}
		fmt.Println()
	}
}
