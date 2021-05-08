package config

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"os"
)

type Config struct {
	Common Common `toml:"common"`
}
type Common struct {
	WithSeconds bool     `toml:"with_seconds"`
	Schedule    string   `toml:"schedule"`
	Command     string   `toml:"command"`
	Args        []string `toml:"args"`
}

func GetConfigFromFile(filePath string) (*Config, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	if err = toml.Unmarshal(b, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
