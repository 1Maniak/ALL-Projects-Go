package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"setup_wizard/internal/broker"
	"setup_wizard/internal/config"
	"setup_wizard/internal/menu"
	"setup_wizard/internal/stages"
	"setup_wizard/internal/state"
)

// Очистка экрана
func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Основная функция
func main() {
	ClearScreen()
	log.SetFlags(0)

	// Загружаем конфиг
	cfg := config.LoadConfig()

	// Вывод информации о подключении
	if cfg.Warning != "" {
		fmt.Println(cfg.Warning)
	}
	fmt.Println(cfg.BrokerInfo)
	fmt.Println()

	// Передаем сообщение в меню
	menu.WarningMessage = cfg.Warning
	menu.BrokerInfoMessage = cfg.BrokerInfo
	
	// Подключаемся к брокеру
	b := broker.NewClient(cfg.BrokerURL)
	if err := b.Connect(); err != nil {
		log.Fatalf("Ошибка подключения к брокеру: %v", err)
	}
	defer b.Close()

	// Загружаем состояние
	state := state.LoadState()

	// Создание этапов
	stageList := []stages.Stage{
		stages.NewAdminStage(b, state),
		// Добавление других этапов...
	}

	// Показываем меню
	menu := menu.NewMenu(stageList)
	if err := menu.Show(); err != nil {
		log.Printf("Ошибка: %v", err)
		os.Exit(1)
	}
}
