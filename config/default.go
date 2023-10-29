package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	MongoDBURL             string        `mapstructure:"MONGODB_URL"`
	RedisURI               string        `mapstructure:"REDIS_URL"`
	RedisPassword          string        `mapstructure:"REDIS_PASS"`
	PORT                   string        `mapstructure:"PORT"`
	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge     int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`
}

func LoadConfig(path string) (c *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()
	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&c)
	return
}
