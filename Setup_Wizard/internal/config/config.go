package config

import (
	"flag"
	"fmt"
)

// Хранит настройки
type Config struct {
	BrokerURL  string
	Warning    string
	BrokerInfo string
}

// Загружает конфигурацию из CLI - флагов
func LoadConfig() Config {
	var brokerURL string

	// Парсим флаг
	flag.StringVar(&brokerURL, "broker-url", "", "URL для подключения к брокеру сообщений")
	flag.Parse()

	var warning string

	// Если не указан, значение по умолчанию
	if brokerURL == "" {
		brokerURL = "amqp://admin@admin@localhost:5672"
		warning = "Внимание: --broker-url не указан, используется значение по умолчанию!"
	}

	brokerInfo := fmt.Sprintf("Подключение к брокеру по адресу: %s", brokerURL)

	return Config{
		BrokerURL:  brokerURL,
		Warning:    warning,
		BrokerInfo: brokerInfo,
	}
}
