package config

import (
	"strings"

	grpcconfig "pillsreminder/apps/pillsservice/api/grpc/config"
	authconfig "pillsreminder/internal/mobileapi/businesslayer/auth/config"
	commonconfig "pillsreminder/pkg/db/config"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Server ServerConfiguration
	DB     commonconfig.DatalayerConfig
	Grpc   grpcconfig.GrpcConfig
}

type ServerConfiguration struct {
	Port   string
	Mode   string
	Secure authconfig.Secure
}

func SetupConfig() AppConfig {
	var configuration AppConfig

	viper.AddConfigPath("etc")
	viper.AddConfigPath("/src/apps/pillsservice/etc")
	viper.SetConfigName("pillsservice.config.yml")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config file")
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		panic("Error reading config file")
	}

	return configuration
}
