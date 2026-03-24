package modules

import (
	"fmt"
	"time"
)

type DatabaseConfig struct {
	URL             string        `yaml:"url"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
}

func (c *DatabaseConfig) Validate() error {
	if c.URL == "" {
		return fmt.Errorf("database.url is required")
	}
	if c.MaxOpenConns <= 0 {
		return fmt.Errorf("database.max_open_conns must be positive")
	}
	if c.MaxIdleConns <= 0 {
		return fmt.Errorf("database.max_idle_conns must be positive")
	}
	if c.MaxIdleConns > c.MaxOpenConns {
		return fmt.Errorf("database.max_idle_conns (%d) must not exceed max_open_conns (%d)", c.MaxIdleConns, c.MaxOpenConns)
	}
	if c.ConnMaxLifetime <= 0 {
		return fmt.Errorf("database.conn_max_lifetime must be positive")
	}
	return nil
}

func (c *DatabaseConfig) ApplyDefaults() {
	if c.URL == "" {
		c.URL = "postgres://postgres:postgres@localhost:5432/app?sslmode=disable"
	}
	if c.MaxOpenConns == 0 {
		c.MaxOpenConns = 25
	}
	if c.MaxIdleConns == 0 {
		c.MaxIdleConns = 5
	}
	if c.ConnMaxLifetime == 0 {
		c.ConnMaxLifetime = 5 * time.Minute
	}
}
