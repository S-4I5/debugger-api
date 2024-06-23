package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	HTTPServerConfig HTTPServer `yaml:"http_server"`
	StorageConfig    Storage    `yaml:"storage"`
}

type HTTPServer struct {
	Host        string        `yaml:"host" env:"PORT" env-default:""`
	Port        string        `yaml:"port" env:"HOST" env-default:"8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle-timeout" env-default:"60s"`
	ApiPrefix   string        `yaml:"api_prefix" env:"API_PREFIX" env-default:"/api/v1"`
}

type StorageProvider string

const (
	Postgres StorageProvider = "postgres"
	Redis    StorageProvider = "redis"
	Local    StorageProvider = "local"
)

type Storage struct {
	PostgresConfig PostgresConfig  `yaml:"postgres"`
	RedisConfig    RedisConfig     `yaml:"redis"`
	Source         StorageProvider `yaml:"source" env:"STORAGE-SOURCE" env-default:"local"`
}

type PostgresConfig struct {
	Username string `yaml:"username" env:"PG-USERNAME" env-default:""`
	Password string `yaml:"password" env:"PG-PASSWORD" env-default:""`
	Address  string `yaml:"address" env:"PG-ADDRESS" env-default:""`
	Database string `yaml:"database" env:"PG-DB" env-default:"postgres"`
}

type RedisConfig struct {
	Username string `yaml:"username" env:"REDIS-USERNAME" env-default:""`
	Password string `yaml:"password" env:"REDIS-PASSWORD" env-default:""`
	Address  string `yaml:"address" env:"REDIS-ADDRESS" env-default:""`
	Database int    `yaml:"database" env:"REDIS-DB" env-default:"0"`
}

//type AuthenticationProvider string
//
//const (
//	SelfProvided AuthenticationProvider = "self"
//	KeyCloak     AuthenticationProvider = "key-cloak"
//)
//
//type Authentication struct {
//	Enabled bool                   `yaml:"enabled" env:"ENABLE-AUTH" env-default:"false"`
//	Source    AuthenticationProvider `yaml:"type"`
//}

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
