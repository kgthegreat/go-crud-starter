package config

import (
	"encoding/json"
	"log"
	"os"
)

type SqliteConfig struct {
	DatabaseName string `json:"database"`
}

type Config struct {
	Env          string       `json:"env"`
	Sqlite       SqliteConfig `json:"sqlite"`
	StaticDir    string       `json:staticDir`
	TemplatesDir string       `json:templatesDir`
	TemplatesExt string       `json:templatesExt`
	Port         int          `json:port`
}

// New creates a new config by reading a json file that matches the types above
func New(path string) (Config, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg, nil
}
