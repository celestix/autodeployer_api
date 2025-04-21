package config

import (
	"io"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	TokenExpirationTime = time.Hour * 24
)

type _Config struct {
	Port                uint   `yaml:"port"`
	Debug               bool   `yaml:"debug"`
	Db_Uri              string `yaml:"db_uri"`
	SecretKey           string `yaml:"secret_key"`
	GhOauthClientId     string `yaml:"gh_oauth_client_id"`
	GhOauthClientSecret string `yaml:"gh_oauth_client_secret"`
	DataDirectory       string `yaml:"data_directory,omitempty"`
}

var Data _Config

func Load() error {
	file, err := os.Open("config.yaml")
	if err != nil {
		return err
	}
	buf, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(buf, &Data)
	if err != nil {
		return err
	}
	if Data.DataDirectory != "" {
		return nil
	}
	setupDataDirectory()
	return nil
}

func setupDataDirectory() {
	const DD = "DATA_DIRECTORY"
	if os.Getenv(DD) != "" {
		Data.DataDirectory = os.Getenv(DD)
		return
	}
	if os.Getenv("DOCKER_ENV") != "" {
		Data.DataDirectory = "/data"
		return
	}
	if _, err := os.Stat(".dockernev"); err == nil {
		Data.DataDirectory = "/data"
		return
	}
	Data.DataDirectory = "./data"
	return
}
