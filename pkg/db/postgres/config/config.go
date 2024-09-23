package config

type PqConfig struct {
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         string
	MaxLifetime  int `mapstructure:"max_life_time"`
	MaxOpenConns int `mapstructure:"max_open_conns"`
	MaxIdleConns int `mapstructure:"max_idle_conns"`
}
