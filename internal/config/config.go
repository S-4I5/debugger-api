package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env" env-required:"true"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Host        string        `yaml:"host" env:"PORT" env-default:""`
	Port        string        `yaml:"port" env:"HOST" env-default:"8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle-timeout" env-default:"60s"`
	BasePath    string        `yaml:"base_path" env:"BASE_PATH" env-default:"/api/v1"`
}

func MustLoad(configPath string) Config {

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("Cannot find config file")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal("Error while reading config")
	}

	return cfg
}
