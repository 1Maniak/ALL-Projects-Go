package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	seconds   int
	isRunning bool
	timeLabel *widget.Label
)

func formatTime(total int) string {
	hours := total / 3600
	minutes := (total % 3600) / 60
	seconds := total % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func updateLabel() {
	if timeLabel != nil {
		fyne.Do(func() {
			timeLabel.SetText(formatTime(seconds))
		})
	}
}

func startTimer() {
	if isRunning {
		fmt.Println("Таймер уже идет!")
		return
	}
	fmt.Println("Запускаю таймер!")
	isRunning = true

	go func() {
		for isRunning {
			time.Sleep(1 * time.Second)
			seconds++
			fmt.Println("прошло секунд: ", seconds)
			updateLabel()
		}
	}()
}

func stopTimer() {
	fmt.Println("Останавливаю таймер!")
	isRunning = false
}

func resetTimer() {
	isRunning = false
	seconds = 0
	updateLabel()
}

func main() {
	// Создаем приложение
	a := app.New()

	// Создаем окно
	w := a.NewWindow("Секундомер v0.1.0")

	// Размер окна
	w.Resize(fyne.NewSize(350,200))

	// Создаем метку времени
	timeLabel = widget.NewLabel("00:00:00")
	timeLabel.Alignment = fyne.TextAlignCenter

	// Создаем кнопки
	startBtn := widget.NewButton("Старт", startTimer)
	stopBtn := widget.NewButton("Стоп", stopTimer)
	resetBtn := widget.NewButton("Сброс", resetTimer)
	closeBtn := widget.NewButton("Закрыть", func() {
		a.Quit()
	})

	// Размещаем элементы
	buttons := container.NewHBox(
		container.NewPadded(startBtn),
		container.NewPadded(stopBtn),
		container.NewPadded(resetBtn),
	)

	 content := container.NewVBox(
		layout.NewSpacer(),
        container.NewCenter(timeLabel),
        container.NewCenter(buttons),
        container.NewCenter(closeBtn),
		layout.NewSpacer(),
    )

	// Устанавливаем содержимое в окно
	w.SetContent(content)

	// Показываем окно и запускаем
	w.ShowAndRun()
}
