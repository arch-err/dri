package config

import (
	"fmt"
	"os"
)

type Config struct {
	Server string
	Token  string
}

func Load() (Config, error) {
	server := os.Getenv("DRONE_SERVER")
	if server == "" {
		return Config{}, fmt.Errorf("DRONE_SERVER environment variable is not set")
	}

	token := os.Getenv("DRONE_TOKEN")
	if token == "" {
		return Config{}, fmt.Errorf("DRONE_TOKEN environment variable is not set")
	}

	return Config{Server: server, Token: token}, nil
}
