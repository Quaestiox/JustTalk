package util

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	DBDrive             string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	HTTPServerAddress   string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress   string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	RedisAddress        string        `mapstructure:"REDIS_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	MQ_URL              string        `mapstructure:"MQ_URL"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
