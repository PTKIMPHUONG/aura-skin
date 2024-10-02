package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	Neo4j     Neo4jConfig
	SecretKey string
}

type Neo4jConfig struct {
	URI      string
	Username string
	Password string
}

func Instance() (*Config, error) {
	var err error
	once.Do(func() {
		viper.SetConfigName("config")                 // Tên file là "config"
		viper.SetConfigType("yaml")                   // Định dạng YAML
		viper.AddConfigPath("./internal/configs/dev") // Đường dẫn đến folder chứa file config

		err = viper.ReadInConfig()
		if err != nil {
			err = log.Output(2, "unable to read config: "+err.Error())
			return
		}

		instance = &Config{}
		err = viper.Unmarshal(instance)
		if err != nil {
			err = log.Output(2, "unable to decode into struct: "+err.Error())
		}
	})

	return instance, err
}

// GetSecretKey trả về SecretKey
func (cfg *Config) GetSecretKey() string {
	return cfg.SecretKey
}
