package config

import (
	"fmt"
	"os"
)

type DbConfig struct {
	Url      string
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	Name     string `env:"DB_NAME"`
}
type ServerConfig struct {
	Port string `env:"SERVER_PORT"`
	Host string `env:"SERVER_HOST"`
}
type Config struct {
	Db     DbConfig
	Server ServerConfig
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	cfg.Db.Host = os.Getenv("DB_HOST")
	cfg.Db.User = os.Getenv("DB_USER")
	cfg.Db.Password = os.Getenv("DB_PASSWORD")
	cfg.Db.Name = os.Getenv("DB_NAME")
	cfg.Db.Port = os.Getenv("DB_PORT")
	cfg.Db.Url = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.Host,
		cfg.Db.Port,
		cfg.Db.Name)

	cfg.Server.Port = os.Getenv("SERVER_PORT")
	cfg.Server.Host = os.Getenv("SERVER_HOST")
	return cfg, nil
}
