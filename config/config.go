package config

import (
	"strings"
	"time"

	"github.com/golovanevvs/vortex/rabbitmq"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	v              *viper.Viper
	ServerConfig   serverConfig
	LoggerConfig   loggerConfig
	RabbitMQConfig rabbitmq.Config
}

type serverConfig struct {
	Addr string
}

type loggerConfig struct {
	LogLevel string
}

func New() *Config {

	v := viper.New()

	return &Config{v: v}
}

func (c *Config) Load(pathConfigFile string, pathEnvFile string, envPrefix string) error {
	godotenv.Load(pathEnvFile)

	c.v.AutomaticEnv()

	if envPrefix != "" {
		c.v.SetEnvPrefix(envPrefix)
	}

	c.v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	c.v.SetConfigFile(pathConfigFile)
	err := c.v.ReadInConfig()
	if err != nil {
		return err
	}

	c.ServerConfig.Addr = c.GetString("server.addr")

	c.LoggerConfig.LogLevel = c.GetString("logging.level")

	c.RabbitMQConfig.URL = c.GetString("rabbitmq.url")
	c.RabbitMQConfig.ReconnectDelay = time.Duration(c.GetInt("rabbitmq.reconnect_delay"))
	c.RabbitMQConfig.MaxReconnect = c.GetInt("rabbitmq.max_reconnect")
	c.RabbitMQConfig.Exchange = c.GetString("rabbitmq.exchange")
	c.RabbitMQConfig.ExchangeType = c.GetString("rabbitmq.exchange_type")
	c.RabbitMQConfig.Queue = c.GetString("rabbitmq.queue")
	c.RabbitMQConfig.RoutingKey = c.GetString("rabbitmq.routing_key")
	c.RabbitMQConfig.Durable = c.GetBool("rabbitmq.durable")
	c.RabbitMQConfig.AutoDelete = c.GetBool("rabbitmq.auto_delete")
	c.RabbitMQConfig.Internal = c.GetBool("rabbitmq.internal")
	c.RabbitMQConfig.Exclusive = c.GetBool("rabbitmq.exclusive")
	c.RabbitMQConfig.NoWait = c.GetBool("rabbitmq.no_wait")
	c.RabbitMQConfig.PrefetchCount = c.GetInt("rabbitmq.prefetch_count")
	c.RabbitMQConfig.PrefetchSize = c.GetInt("rabbitmq.prefetch_size")
	c.RabbitMQConfig.GlobalPrefetch = c.GetBool("rabbitmq.global_prefetch")
	c.RabbitMQConfig.Mandatory = c.GetBool("rabbitmq.mandatory")
	c.RabbitMQConfig.Immediate = c.GetBool("rabbitmq.immediate")

	return nil
}

func (c *Config) GetString(key string) string {
	return c.v.GetString(key)
}

func (c *Config) GetInt(key string) int {
	return c.v.GetInt(key)
}

func (c *Config) GetBool(key string) bool {
	return c.v.GetBool(key)
}
