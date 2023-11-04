package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	EmailFrom              string        `mapstructure:"EMAIL_FROM"`
	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	SMTPUser               string        `mapstructure:"SMTP_USER"`
	PORT                   string        `mapstructure:"PORT"`
	SMTPPass               string        `mapstructure:"SMTP_PASS"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	SMTPHost               string        `mapstructure:"SMTP_HOST"`
	RedisPassword          string        `mapstructure:"REDIS_PASS"`
	RedisURI               string        `mapstructure:"REDIS_URL"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	MongoDBURL             string        `mapstructure:"MONGODB_URL"`
	RefreshTokenMaxAge     int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`
	AccessTokenMaxAge      int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	SMTPPort               int           `mapstructure:"SMTP_PORT"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
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
