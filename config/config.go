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
