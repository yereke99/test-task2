package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Url string `yaml:"URL"`

	User     string `yaml:"USER"`
	Password string `yaml:"PASSWORD"`
	Host     string `yaml:"HOST"`
	Port     string `yaml:"PORT"`
	Database string `yaml:"DATABASE"`
}

func NewConfig(fileName string) (*Config, error) {

	cfg := new(Config)

	if err := loadEnv(fileName); err != nil {
		panic(err)
	}

	cfg.Url = os.Getenv("URL")
	cfg.User = os.Getenv("USER")
	cfg.Password = os.Getenv("PASSWORD")
	cfg.Host = os.Getenv("HOST")
	cfg.Port = os.Getenv("PORT")
	cfg.Database = os.Getenv("DATABASE")

	return cfg, nil
}

func loadEnv(fileName string) error {

	if err := godotenv.Load(fileName); err != nil {
		return err
	}
	return nil
}
