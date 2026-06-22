package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Загружает задачи из файла JSON
func loadTasks() ([]Task, error) {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON:", err)
		return nil, err
	}
	return tasks, nil
}

// Сохраняет задачи в файл
func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
} 

// Возращает следующий id для новой задачи
func getNextID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}
	maxID := 0 
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}
