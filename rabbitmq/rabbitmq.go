package rabbitmq

import (
	"errors"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

// RabbitMQ клиент
type Client struct {
	logger  *zerolog.Logger
	conn    *amqp.Connection
	channel *amqp.Channel
	config  Config
}

type Config struct {
	URL                   string
	ReconnectDelaySeconds int
	MaxReconnect          int
	Exchange              string
	ExchangeType          string
	Queue                 string
	RoutingKey            string
	Durable               bool
	AutoDelete            bool
	Internal              bool
	Exclusive             bool
	NoWait                bool
	PrefetchCount         int
	PrefetchSize          int
	GlobalPrefetch        bool
	Mandatory             bool
	Immediate             bool
	DelayedExchange       string
	DelayedQueue          string
}

// Конструктор нового клиента
func New(config Config, logger *zerolog.Logger) (*Client, error) {
	client := &Client{
		logger: logger,
		config: config}

	if err := client.connect(); err != nil {
		return nil, err
	}

	if err := client.setup(); err != nil {
		return nil, err
	}

	go client.reconnectListener()

	return client, nil
}

// Подключение к RabbitMQ
func (c *Client) connect() error {
	conn, err := amqp.DialConfig(c.config.URL, amqp.Config{
		Dial: amqp.DefaultDial(5 * time.Second),
	})
	if err != nil {
		return err
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return err
	}

	if err := channel.Qos(
		c.config.PrefetchCount,
		c.config.PrefetchSize,
		c.config.GlobalPrefetch,
	); err != nil {
		conn.Close()
		return err
	}

	c.conn = conn
	c.channel = channel

	return nil
}

// Настройка exchange и queue
func (c *Client) setup() error {
	if c.config.Exchange != "" {
		if err := c.channel.ExchangeDeclare(
			c.config.Exchange,
			c.config.ExchangeType,
			c.config.Durable,
			c.config.AutoDelete,
			c.config.Internal,
			c.config.NoWait,
			nil,
		); err != nil {
			return err
		}
	}

	if c.config.Queue != "" {
		_, err := c.channel.QueueDeclare(
			c.config.Queue,
			c.config.Durable,
			c.config.AutoDelete,
			c.config.Exclusive,
			c.config.NoWait,
			nil,
		)
		if err != nil {
			return err
		}

		if c.config.Exchange != "" && c.config.RoutingKey != "" {
			if err := c.channel.QueueBind(
				c.config.Queue,
				c.config.RoutingKey,
				c.config.Exchange,
				c.config.NoWait,
				nil,
			); err != nil {
				return err
			}
		}

		if c.config.DelayedExchange != "" {
			args := amqp.Table{"x-delayed-type": "direct"}
			if err := c.channel.ExchangeDeclare(
				c.config.DelayedExchange,
				"x-delayed-message",
				c.config.Durable,
				c.config.AutoDelete,
				c.config.Internal,
				c.config.NoWait,
				args,
			); err != nil {
				return fmt.Errorf("failed to declare delayed exchange: %w", err)
			}

			if c.config.DelayedQueue != "" {
				_, err := c.channel.QueueDeclare(
					c.config.DelayedQueue,
					c.config.Durable,
					c.config.AutoDelete,
					c.config.Exclusive,
					c.config.NoWait,
					nil,
				)
				if err != nil {
					return fmt.Errorf("failed to declare delayed queue: %w", err)
				}

				if err := c.channel.QueueBind(
					c.config.DelayedQueue,
					c.config.RoutingKey,
					c.config.DelayedExchange,
					c.config.NoWait,
					nil,
				); err != nil {
					return fmt.Errorf("failed to bind delayed queue: %w", err)
				}
			}
		}
	}

	return nil
}

// Слушатель реконнектов
func (c *Client) reconnectListener() {
	for {
		reason := <-c.conn.NotifyClose(make(chan *amqp.Error))
		if reason == nil {
			return
		}

		c.logger.Warn().Msgf("Connection closed: %v. Reconnecting...", reason)

		var err error
		for i := 0; i < c.config.MaxReconnect; i++ {
			if err = c.connect(); err == nil {
				c.logger.Info().Msg("Reconnected successfully")
				break
			}

			time.Sleep(time.Duration(c.config.ReconnectDelaySeconds) * time.Second)
		}

		if err != nil {
			c.logger.Error().Err(err).Msgf("Failed to reconnect after %d attempts", c.config.MaxReconnect)
			return
		}
	}
}

// Публикация сообщения
func (c *Client) Publish(body []byte, headers amqp.Table) error {
	if c.channel == nil {
		return errors.New("channel not initialized")
	}

	return c.channel.Publish(
		c.config.Exchange,
		c.config.RoutingKey,
		c.config.Mandatory,
		c.config.Immediate,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Headers:     headers,
		},
	)
}

func (c *Client) PublishDelayed(body []byte, headers amqp.Table, delay time.Duration) error {
	if c.channel == nil {
		return errors.New("channel not initialized")
	}
	if c.config.DelayedExchange == "" {
		return errors.New("delayed exchange not configured")
	}

	// Добавляем заголовок x-delay с указанием задержки в миллисекундах
	if headers == nil {
		headers = make(amqp.Table)
	}
	headers["x-delay"] = int(delay.Milliseconds())

	return c.channel.Publish(
		c.config.DelayedExchange,
		c.config.RoutingKey,
		c.config.Mandatory,
		c.config.Immediate,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Headers:     headers,
		},
	)
}

// Подписка на сообщения
func (c *Client) Consume(handler func([]byte) error) error {
	if c.channel == nil || c.config.Queue == "" {
		return errors.New("channel or queue not configured")
	}

	msgs, err := c.channel.Consume(
		c.config.Queue,
		"",
		false,
		c.config.Exclusive,
		false,
		c.config.NoWait,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			if err := handler(msg.Body); err == nil {
				msg.Ack(false)
			} else {
				msg.Nack(false, true)
			}
		}
	}()

	return nil
}

// Закрытие соединения
func (c *Client) Close() error {
	if c.channel != nil {
		if err := c.channel.Close(); err != nil {
			return err
		}
	}
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
