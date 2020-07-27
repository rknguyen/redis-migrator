package config

import (
	"log"

	"github.com/spf13/viper"
)

type RedisConnection struct {
	Host          string `yaml:"host" mapstructure:"host"`
	Port          int    `yaml:"port" mapstructure:"port"`
	Password      string `yaml:"password" mapstructure:"password"`
	DatabaseIndex int    `yaml:"database_index" mapstructure:"database_index"`
}

type Config struct {
	Src       RedisConnection `yaml:"old" mapstructure:"old"`
	Dst       RedisConnection `yaml:"new" mapstructure:"new"`
	ThreadNum int             `yaml:"thread_num" mapstructure:"thread_num"`
}

func LoadConfig() *Config {
	var cfg = &Config{}
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config/config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to read config", err)
	}

	viper.AutomaticEnv()

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("Failed to unmarshal config", err)
	}

	log.Println("Config loaded")
	return cfg
}
