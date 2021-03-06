package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Server  *Server
	Indexer *Indexer
}

type Server struct {
	Port int
}

type Indexer struct {
	HTTP *HTTPIndexer
}

type HTTPIndexer struct {
	Addr string
}

func Read(path string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(file).Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}

func GetDefaultConfig() *Config {
	c := &Config{
		Server: &Server{
			Port: 8080,
		},
		Indexer: &Indexer{
			HTTP: &HTTPIndexer{
				Addr: ":8081",
			},
		},
	}
	return c
}
