package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name      string `mapstructure:"name"`
		Port      int    `mapstructure:"port"`
		Env       string `mapstructure:"env"`
		JWTSecret string `mapstructure:"jwt_secret"`
	} `mapstructure:"app"`

	DB struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
		SSLMode  string `mapstructure:"sslmode"`
	} `mapstructure:"db"`

	Redis struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"redis"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	log.Println("Config loaded successfully")
	return &cfg, nil
}
