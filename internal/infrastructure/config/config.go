package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type (
	Config struct {
		Env        string            `yaml:"env" env-default:"dev"`
		Server     AuthServiceServer `yaml:"http_server"`
		Postgres   PostgreSQL        `yaml:"postgres"`
		Secret     Secret            `yaml:"secret"`
		SuperAdmin SuperAdmin        `yaml:"super_admin"`
		Path       Path              `yaml:"path"`
	}

	AuthServiceServer struct {
		Address     string        `yaml:"address" env-default:"localhost:5000"`
		Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
		IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	}

	PostgreSQL struct {
		Host     string `yaml:"host" env-default:"localhost"`
		Port     int    `yaml:"port" env-default:"5121"`
		User     string `yaml:"user" env-default:"postgres"`
		Password string `yaml:"password" env-default:"root"`
		DbName   string `yaml:"dbname" env-default:"database"`
	}

	Secret struct {
		ApiKey           string `yaml:"api_key" env-defaul:"super-puper-secret-key"`
		AccessTokenTime  string `yaml:"access_token_time" env-default:"10m"`
		RefreshTokenTime string `yaml:"refresh_token_time" env-default:"1h"`
	}

	SuperAdmin struct {
		Login    string `yaml:"login"`
		Password string `yaml:"password"`
		Email    string `yaml:"email"`
	}

	Path struct {
		LogFile string `yaml:"log_file"`
	}
)

var instance *Config = nil

func MustLoad() *Config {
	if instance != nil {
		return instance
	}
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	instance = &cfg

	return instance
}
