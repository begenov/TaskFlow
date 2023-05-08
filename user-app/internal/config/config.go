package config

import "time"

type (
	Config struct {
		JWT JWTConfig
	}

	JWTConfig struct {
		AccessTokenTTL  time.Duration
		RefreshTokenTTL time.Duration
		SigninKey       string
	}
)

func NewConfig() {
}
