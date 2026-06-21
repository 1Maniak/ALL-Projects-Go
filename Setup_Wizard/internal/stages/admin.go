package stages

import (
	"fmt"
	"time"

	"setup_wizard/internal/broker"
	"setup_wizard/internal/state"

	"github.com/AlecAivazis/survey/v2"
)

// Этап настройки админа
type AdminStage struct {
	broker *broker.Client
	state  *state.State
}

// Создает новый этап
func NewAdminStage(b *broker.Client, s *state.State) *AdminStage {
	return &AdminStage{
		broker: b,
		state:  s,
	}
}

// Возвращает название этапа
func (s *AdminStage) Name() string {
	return "Настройка администратора"
}

// Проверяем, выполнен ли этап
func (s *AdminStage) IsCompleted() (bool, error) {
	// Проверяем состояние
	if s.state.IsCompleted("admin") {
		return true, nil
	}

	// Проверяем через брокер
	exists, err := s.broker.IsAdminExists()
	if err != nil {
		return false, err
	}

	if exists {
		s.state.MarkCompleted("admin")
		return true, nil
	}
	return false, err
}

// Выполняем этапы
func (s *AdminStage) Run() error {

	// Проверяем не выполнен ли уже этап
	completed, err := s.IsCompleted()
	if err != nil {
		return err
	}
	if completed {
		fmt.Println("Администратор уже настроен")
		return nil
	}

	// Запрашиваем логин c валидацией через survey
	var login string
	prompt := &survey.Input{
		Message: "Введите логин администратора:",
	}

	err = survey.AskOne(prompt, &login, survey.WithValidator(func(val interface{}) error {
		str, ok := val.(string) 
		if !ok {
			return fmt.Errorf("Введите строку")
		}
		return broker.ValidateLogin(str)
	}))
	if err != nil {
		return err
	}

	for {
		// Валидация логина
		if err := broker.ValidateLogin(login); err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			continue
		}
		break
	}

	// log - Создаем администратора
	// log.Printf("Создание администратора: %s", login)

	// Индикатор загрузки
	fmt.Print("Отправка запроса в брокер")
	for i := 0; i < 4; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println("\nГотово!")

	if err := s.broker.CreateAdmin(login); err != nil {
		return fmt.Errorf("ошибка создания администратора: %w", err)
	}

	fmt.Println("Нажмите Enter для продолжения...")
	fmt.Scanln()

	// Отмечаем этап как выполненный
	s.state.MarkCompleted("admin")
	fmt.Println("Администратор успешно создан!")

	return nil
}

// Возвращает статус этапа
func (s *AdminStage) Status() string {
	completed, _ := s.IsCompleted()
	if completed {
		return "[X]"
	}
	return "[ ]"
}
