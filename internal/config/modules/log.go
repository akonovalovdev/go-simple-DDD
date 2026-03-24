package modules

import "fmt"

type LogConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

func (c *LogConfig) Validate() error {
	validLevels := map[string]bool{"debug": true, "info": true, "warn": true, "error": true}
	if !validLevels[c.Level] {
		return fmt.Errorf("log.level must be one of [debug, info, warn, error], got %q", c.Level)
	}
	validFormats := map[string]bool{"json": true, "text": true}
	if !validFormats[c.Format] {
		return fmt.Errorf("log.format must be one of [json, text], got %q", c.Format)
	}
	return nil
}

func (c *LogConfig) ApplyDefaults() {
	if c.Level == "" {
		c.Level = "info"
	}
	if c.Format == "" {
		c.Format = "json"
	}
}
