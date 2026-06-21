package state

import (
	"encoding/json"
	"os"
)

const stateFile = "state.json"

// Хранит состояние выполненных этапов
type State struct {
	CompletedStages map[string]bool `json:"completed_stages"`
}

// Загружаем состояние из файла
func LoadState() *State {
	data, err := os.ReadFile(stateFile)
	if err != nil {
		return &State{
			CompletedStages: make(map[string]bool),
		}
	}

	var state State
	if err := json.Unmarshal(data, &state); err != nil {
		return &State{
			CompletedStages: make(map[string]bool),
		}
	}
	return &state
}

// Сохраняем состояние в файл
func (s *State) Save() error {
	data, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(stateFile, data, 0644)
}

// Проверяем, выполнен ли этап
func (s *State) IsCompleted(stageName string) bool {
	return s.CompletedStages[stageName]
}

// Отмечаем выполненный этап как выполненный
func (s *State) MarkCompleted(stageName string) {
	s.CompletedStages[stageName] = true
	s.Save()
}
