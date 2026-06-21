package broker

import (
	"fmt"
	"log"
)

// Клиент для работы с брокером
type Client struct {
	url string
}

// Создает нового клиента
func NewClient(url string) *Client {
	return &Client{
		url: url,
	}
}

// Подключение к брокеру
func (c *Client) Connect() error {
	log.Printf("Подключение к брокеру по адресу: %s", c.url)
	// ......................!!!
	return nil
}

// Проверяет существует ли админ
func (c *Client) IsAdminExists() (bool, error) {
	// .....................!!!
	return false, nil
}

// Создает администратора
func (c *Client) CreateAdmin(login string) error {
	// ......................!!!
	// log.Printf("Создание администратора: %s", login)
	return nil
}

// Закрыват соединение
func (c *Client) Close() error {
	log.Println("Закрытие соединения с брокером")
	return nil
}

// Проверяем логин на валидность
func ValidateLogin(login string) error {
	if login == "" {
		return fmt.Errorf("логин не может быть пустым")
	}
	if len(login) < 3 {
		return fmt.Errorf("логин должен содержать минимум 3 символа")
	}
	return nil
}
