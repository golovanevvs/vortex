package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	v *viper.Viper
}

func New() *Config {
	v := viper.New()

	v.AutomaticEnv()

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return &Config{v: v}
}

func (c *Config) SetEnvPrefix(prefix string) {
	c.v.SetEnvPrefix(prefix)
}

func (c *Config) Load(path string) error {
	c.v.SetConfigFile(path)

	return c.v.ReadInConfig()
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
