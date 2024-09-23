package config

import (
	"strings"

	grpcconfig "pillsreminder/apps/reminderservice/api/grpc/config"
	"pillsreminder/pkg/config"
	commonconfig "pillsreminder/pkg/db/config"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Server           ServerConfiguration
	DB               commonconfig.DatalayerConfig
	Grpc             grpcconfig.GrpcConfig
	ExternalServices ExternalServices
}

type ServerConfiguration struct {
	Port string
	Mode string
}

type ExternalServices struct {
	ReminderService config.ExternalService `json:"mobileapi"`
}

func SetupConfig() AppConfig {
	var configuration AppConfig

	viper.AddConfigPath("etc")
	viper.AddConfigPath("/src/apps/reminderservice/etc")
	viper.SetConfigName("reminderservice.config.yml")
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
