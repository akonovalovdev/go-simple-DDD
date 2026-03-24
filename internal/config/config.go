package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/akonovalovdev/go-simple-DDD/internal/config/modules"
)

// Config — единый конфиг приложения.
// Каждое поле — отдельный модульный конфиг из internal/config/modules.
// Внешний API не меняется: cfg.Server, cfg.Database, cfg.Log и т.д.
type Config struct {
	Server   modules.ServerConfig   `yaml:"server"`
	Database modules.DatabaseConfig `yaml:"database"`
	Log      modules.LogConfig      `yaml:"log"`
}

// Load читает yaml-файл, проставляет defaults, валидирует и возвращает готовый конфиг.
// Pipeline: loadFromFile → setDefaults → validate.
func Load(path string) (*Config, error) {
	cfg := &Config{}

	if err := cfg.loadFromFile(path); err != nil {
		return nil, fmt.Errorf("failed to load config from file: %w", err)
	}

	cfg.setDefaults()

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return cfg, nil
}

func (c *Config) loadFromFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, c)
}

// setDefaults делегирует проставление defaults каждому модульному конфигу.
func (c *Config) setDefaults() {
	c.Server.ApplyDefaults()
	c.Database.ApplyDefaults()
	c.Log.ApplyDefaults()
}

// validate делегирует валидацию каждому модульному конфигу.
func (c *Config) validate() error {
	if err := c.Server.Validate(); err != nil {
		return err
	}
	if err := c.Database.Validate(); err != nil {
		return err
	}
	if err := c.Log.Validate(); err != nil {
		return err
	}
	return nil
}
