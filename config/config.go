package config

import (
	"io/ioutil"
	"os"

	"github.com/pelletier/go-toml"
)

type Config struct {
	Common Common `toml:"common"`
}
type Common struct {
	Schedule string   `toml:"schedule"`
	Command  string   `toml:"command"`
	Args     []string `toml:"args"`
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
