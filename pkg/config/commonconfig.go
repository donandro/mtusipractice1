package config

import "time"

type ExternalService struct {
	Address   string        `json:"address"`
	Timeout   time.Duration `json:"timeout"`
	RateLimit int           `json:"ratelimit"`
}
