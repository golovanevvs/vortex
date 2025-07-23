package rabbitmq

import (
	"context"
	"errors"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

// RabbitMQ client
type Client struct {
	cancel  context.CancelFunc
	logger  zerolog.Logger
	conn    *amqp.Connection
	channel *amqp.Channel
	config  Config
}

type Config struct {
	URL                       string
	ReconnectDelaySeconds     int
	MaxReconnect              int
	Exchange                  string
	ExchangeType              string
	Queue                     string
	RoutingKey                string
	Durable                   bool
	AutoDelete                bool
	Internal                  bool
	Exclusive                 bool
	NoWait                    bool
	PrefetchCount             int
	PrefetchSize              int
	GlobalPrefetch            bool
	Mandatory                 bool
	Immediate                 bool
	DelayedExcnahge           string
	DelayedQueue              string
	DelayedExchangeWithPlugin string
	DelayedQueueWithPlugin    string
}

// New client builder
func New(cancelFunc context.CancelFunc, config Config, logger *zerolog.Logger) (*Client, error) {
	client := &Client{
		cancel: cancelFunc,
		logger: logger.With().Str("component", "RabbitMQ").Logger(),
		config: config,
	}

	if err := client.connect(); err != nil {
		logger.Error().Err(err).Msg("Failed to connect")
		return nil, err
	}

	if err := client.setup(); err != nil {
		logger.Error().Err(err).Msg("Failed to client setup")
		return nil, err
	}

	go client.reconnectListener()

	return client, nil
}

// Connecting to RabbitMQ
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

// Setupping exchange and queue
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
			return fmt.Errorf("failed to declare exchange: %w", err)
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
			return fmt.Errorf("failed to declare queue: %w", err)
		}

		if c.config.Exchange != "" && c.config.RoutingKey != "" {
			if err := c.channel.QueueBind(
				c.config.Queue,
				c.config.RoutingKey,
				c.config.Exchange,
				c.config.NoWait,
				nil,
			); err != nil {
				return fmt.Errorf("failed to bind queue: %w", err)
			}
		}

		// delayed

		// rabbitmq-delayed-message-exchange

		if c.config.DelayedExchangeWithPlugin != "" {
			args := amqp.Table{"x-delayed-type": "direct"}
			if err := c.channel.ExchangeDeclare(
				c.config.DelayedExchangeWithPlugin,
				"x-delayed-message",
				c.config.Durable,
				c.config.AutoDelete,
				c.config.Internal,
				c.config.NoWait,
				args,
			); err != nil {
				return fmt.Errorf("failed to declare delayed exchange: %w", err)
			}

			if c.config.DelayedQueueWithPlugin != "" {
				_, err := c.channel.QueueDeclare(
					c.config.DelayedQueueWithPlugin,
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
					c.config.DelayedQueueWithPlugin,
					c.config.RoutingKey,
					c.config.DelayedExchangeWithPlugin,
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

// Reconnect listener
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
			c.cancel()
			return
		}
	}
}

// Publishing
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

// Publishing delayed with plugin
func (c *Client) PublishDelayedWithPlugin(body []byte, delay time.Duration) error {
	if c.channel == nil {
		return errors.New("channel not initialized")
	}
	if c.config.DelayedExchangeWithPlugin == "" {
		return errors.New("delayed exchange not configured")
	}

	headers := make(amqp.Table)
	headers["x-delay"] = int(delay.Milliseconds())

	return c.channel.Publish(
		c.config.DelayedExchangeWithPlugin,
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

// Consuming
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
		c.logger.Error().Err(err).Msg("Failed starts delivering queded messages")
		return err
	}

	go func() {
		for msg := range msgs {
			if err := handler(msg.Body); err == nil {
				msg.Ack(false)
			} else {
				msg.Nack(false, false)
			}
		}
	}()

	return nil
}

// Delayed consuming with plugin
func (c *Client) DelayedConsumeWithPlugin(handler func([]byte) error) error {
	if c.channel == nil || c.config.Queue == "" {
		return errors.New("channel or queue not configured")
	}

	msgs, err := c.channel.Consume(
		c.config.DelayedQueueWithPlugin,
		"",
		false,
		c.config.Exclusive,
		false,
		c.config.NoWait,
		nil,
	)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed starts delivering queded messages")
		return err
	}

	go func() {
		for msg := range msgs {
			if err := handler(msg.Body); err == nil {
				msg.Ack(false)
			} else {
				msg.Nack(false, false)
			}
		}
	}()

	return nil
}

// Connection closing
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
