package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func showMenu() {
	fmt.Println("╔════════════════════════════╗")
	fmt.Println("║    МЕНЕДЖЕР ЗАДАЧ v0.1.0   ║")
	fmt.Println("║         by Maniak          ║")
	fmt.Println("╚════════════════════════════╝")
}

func main() {
	showMenu()
	// Загружаем задачи
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Использование:")
	fmt.Println(" list						- показать все задачи")
	fmt.Println(" add <текст>				- добавить задачу")
	fmt.Println(" done <id>					- отметить задачу выполненной")
	fmt.Println(" delete <id>				- удалить задачу")
	fmt.Println("\nПримеры:")
	fmt.Println(" add Купить хлеб")
	fmt.Println(" done 1")
	fmt.Println(" delete 1")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		args := strings.Fields(input)
		command := args[0]

		switch command {
		case "exit":
			fmt.Println("До свидания!")
			return

		case "list":
			listTasks(tasks)

		case "add":
			if len(args) < 2 {
				fmt.Println("Укажите текст задачи")
				fmt.Println("Пример: add Купить хлеб")
				continue
			}
			title := strings.Join(args[1:], " ")
			tasks = addTask(tasks, title)
			saveTasks(tasks)
			fmt.Printf("Задача добавлена: %s\n", title)

		case "done":
			if len(args) < 2 {
				fmt.Println("Укажите ID задачи")
				fmt.Println("Пример: done 1")
				continue
			}
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("ID должен быть числом")
				continue
			}
			tasks, err = doneTask(tasks, id)
			if err != nil {
				fmt.Println("X", err)
				continue
			}
			saveTasks(tasks)
			fmt.Printf("Задача %d выполнена!\n", id)

		case "delete":
			if len(args) < 2 {
				fmt.Println("Укажите ID задачи")
				fmt.Println("Пример: delete 1")
				continue
			}
			id, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("ID должен быть числом")
				continue
			}
			tasks, err = deleteTask(tasks, id)
			if err != nil {
				fmt.Println("X", err)
				continue
			}
			saveTasks(tasks)
			fmt.Printf("Задача %d удалена!\n", id)

		case "clear":
			tasks = []Task{}
			saveTasks(tasks)
			fmt.Println("Все задачи удалены")

		default:
			fmt.Printf("Неизвестная команда: %s\n", command)
			fmt.Println("Доступные команды: list, add, done, delete, clear, exit")
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("❌ Ошибка чтения ввода:", err)
		}
	}
}

// Выводит список задач
func listTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("Список задач пуст")
		return
	}
	fmt.Println("\nСписок задач:")
	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "x"
		}
		fmt.Printf("  %d. [%s] %s\n", task.ID, status, task.Title)
	}
	fmt.Println()
}

// Добавляет новую задачу
func addTask(tasks []Task, title string) []Task {
	newTask := Task{
		ID:        getNextID(tasks),
		Title:     title,
		Completed: false,
	}
	return append(tasks, newTask)
}

// Отмечает задачу как выполненную
func doneTask(tasks []Task, id int) ([]Task, error) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			return tasks, nil
		}
	}
	return tasks, fmt.Errorf("задача с ID %d не найдена", id)
}

// Удаляет задачу по ID
func deleteTask(tasks []Task, id int) ([]Task, error) {
	for i := range tasks {
		if tasks[i].ID == id {
			return append(tasks[:i], tasks[i+1:]...), nil
		}
	}
	return tasks, fmt.Errorf("задача с ID %d не найдена", id)
}
