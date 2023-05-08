package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/subosito/gotenv"
)

const path = "../.env"

type (
	Config struct {
		JWT jwt_config
	}

	jwt_config struct {
		AccessTokenTTL  time.Duration
		RefreshTokenTTL time.Duration
		SigninKey       string
	}
)

func NewConfig() (*Config, error) {
	err := gotenv.Load(path)
	if err != nil {
		return nil, err
	}

	accessTokenTTL, err := strconv.ParseInt(os.Getenv("ACCESS_TOKEN_TTL"), 10, 64)
	if err != nil {
		log.Fatalf("Error parsing ACCESS_TOKEN_TTL: %v", err)
	}

	refreshTokenTTL, err := strconv.ParseInt(os.Getenv("REFRESH_TOKEN_TTL"), 10, 64)
	if err != nil {
		log.Fatalf("Error parsing REFRESH_TOKEN_TTL: %v", err)
	}

	return &Config{
		JWT: jwt_config{
			AccessTokenTTL:  time.Duration(accessTokenTTL) * time.Minute,
			RefreshTokenTTL: time.Duration(refreshTokenTTL) * time.Minute,
			SigninKey:       os.Getenv("SIGNIN_KEY"),
		},
	}, nil
}
