package config

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	v            *viper.Viper
	AppConfig    appConfig
	LoggerConfig loggerConfig
}

type appConfig struct {
	Addr string
}

type loggerConfig struct {
	LogLevel string
}

func New(envPrefix string) *Config {
	godotenv.Load("../../.env")

	v := viper.New()

	v.AutomaticEnv()

	if envPrefix != "" {
		v.SetEnvPrefix(envPrefix)
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return &Config{v: v}
}

func (c *Config) SetEnvPrefix(prefix string) {
	c.v.SetEnvPrefix(prefix)
}

func (c *Config) Load(path string) error {
	c.v.SetConfigFile(path)
	err := c.v.ReadInConfig()
	if err != nil {
		return err
	}

	c.AppConfig.Addr = c.GetString("server.addr")
	c.LoggerConfig.LogLevel = c.GetString("logging.level")

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
