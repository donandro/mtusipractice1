package config

type GrpcConfig struct {
	Server GrpcServer `mapstructure:"server"`
}

type GrpcServer struct {
	Port int
}
