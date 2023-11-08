package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	SMTPHost               string        `mapstructure:"SMTP_HOST"`
	GRPCServerAddress      string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	EmailFrom              string        `mapstructure:"EMAIL_FROM"`
	SMTPUser               string        `mapstructure:"SMTP_USER"`
	PORT                   string        `mapstructure:"PORT"`
	SMTPPass               string        `mapstructure:"SMTP_PASS"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	Origin                 string        `mapstructure:"ORIGIN"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	RedisURI               string        `mapstructure:"REDIS_URL"`
	MongoDBURL             string        `mapstructure:"MONGODB_URL"`
	RedisPassword          string        `mapstructure:"REDIS_PASS"`
	AccessTokenMaxAge      int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	SMTPPort               int           `mapstructure:"SMTP_PORT"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenMaxAge     int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`
}

func LoadConfig(path string) (c Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()
	if err = viper.ReadInConfig(); err != nil {
		return Config{}, err
	}
	err = viper.Unmarshal(&c)
	return
}
