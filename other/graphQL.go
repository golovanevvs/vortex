package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// Модель данных (структура Message)
type Message struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

// Глобальная переменная для хранения сообщений (упрощённый вариант)
var messages []*Message
var nextID = 1
var mu sync.RWMutex

// Основной класс резольверов
type Resolver struct {
	Message *MessageResolver
}

// Реализация методов для Query-типа
func (r *Resolver) Messages() ([]*Message, error) {
	mu.RLock()
	defer mu.RUnlock()
	return messages, nil
}

// Реализация метода для Mutation-типа
func (r *Resolver) SendMessage(args struct {
	Text   string
	RoomID int
}) bool {
	mu.Lock()
	defer mu.Unlock()

	msg := &Message{
		ID:        strconv.Itoa(nextID),
		Text:      args.Text,
		Timestamp: time.Now(),
	}
	messages = append(messages, msg)
	nextID++
	return true
}

// Резольвер для отдельного типа Message
type MessageResolver struct{}

// Реализация конкретных полей Message
func (mr *MessageResolver) ID(m *Message) string {
	return m.ID
}

func (mr *MessageResolver) Text(m *Message) string {
	return m.Text
}

func (mr *MessageResolver) Timestamp(m *Message) string {
	return m.Timestamp.Format(time.RFC3339)
}

// Inline-схема GraphQL
const schemaStr = `
type Query {
  messages: [Message!]
}

type Mutation {
  sendMessage(text: String!, roomID: Int!): Boolean
}

type Subscription {
  newMessage(roomID: Int!): Message!
}

type Message {
  id: ID!
  text: String!
  timestamp: String!
}
`

// Запуск сервера
func main() {
	// Парсим схему
	parsedSchema, err := graphql.ParseSchema(schemaStr, &Resolver{
		Message: &MessageResolver{},
	})
	if err != nil {
		log.Fatalf("Error parsing schema: %v", err)
	}

	// Настраиваем Relay Handler
	server := &relay.Handler{Schema: parsedSchema}

	// Маршрутизация запросов
	http.Handle("/", server)

	// Запускаем HTTP-сервер
	port := ":8080"
	log.Printf("Starting server on %s...", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}

// JSON сериализация для отображения сообщений
func (m *Message) MarshalJSON() ([]byte, error) {
	type Alias Message
	return json.Marshal(&struct {
		Alias
	}{
		Alias: Alias(*m),
	})
}
