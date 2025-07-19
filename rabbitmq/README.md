# rabbitmq

## Установка

### Docker

```bash
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:4-management
```

## Плагины

### rabbitmq-delayed-message-exchange

[https://github.com/rabbitmq/rabbitmq-delayed-message-exchange](https://github.com/rabbitmq/rabbitmq-delayed-message-exchange)

## Использование

```go
package main

import (
    "log"
    "time"

    "github.com/yourname/rabbitmq"
)

func main() {
    client, err := rabbitmq.New(rabbitmq.Config{
        URL:            "amqp://guest:guest@localhost:5672/",
        Exchange:       "events",
        ExchangeType:   "topic",
        Queue:          "notifications",
        RoutingKey:     "notification.*",
        Durable:        true,
        ReconnectDelay: 2 * time.Second,
        MaxReconnect:   10,
        PrefetchCount:  10,
    })
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    // Публикация
    err = client.Publish([]byte(`{"message": "Hello"}`), nil)
    if err != nil {
        log.Println("Publish error:", err)
    }

    // Подписка
    err = client.Consume(func(body []byte) error {
        log.Println("Received:", string(body))
        return nil
    })
    if err != nil {
        log.Fatal(err)
    }

    // Ожидание сообщений
    select {}
}
```

## Ключевые возможности

- Автоматический реконнект при разрыве соединения
- Гибкая конфигурация через структуру Config
- Поддержка:
  - Публикации сообщений (Publish)
  - Подписки на сообщения (Consume)
  - Подтверждения обработки (Ack/Nack)
- Управление QoS (prefetch count)
- Безопасное закрытие соединений

## Дополнительные улучшения

- Добавить TLS поддержку в Config
- Реализовать подтверждение публикации (Publisher Confirms)
- Добавить метрики (количество сообщений, ошибки и т.д.)
- Реализовать dead-letter queue обработку
