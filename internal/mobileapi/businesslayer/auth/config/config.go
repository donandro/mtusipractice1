package config

import "time"

type Secure struct {
	JwtSecret           string        `mapstructure:"jwt_secret"`
	AccessTokenLifetime time.Duration `mapstructure:"access_token_lifetime"`
}
