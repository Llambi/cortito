package config

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	UserTTL   time.Duration
	UserQuota uint64
	UrlTTL    time.Duration
}

var once = sync.Once{}
var instance *Config

func GetInstance() *Config {
	if instance == nil {
		once.Do(
			func() {
				godotenv.Load()
				userTTL := 1800
				userQuota := 30
				urlTTL := 432000

				if val, exists := os.LookupEnv("USER_TTL"); exists {
					if val, err := strconv.Atoi(val); err == nil {
						userTTL = val
					}
				}

				if val, exists := os.LookupEnv("USER_QUOTA"); exists {
					if val, err := strconv.Atoi(val); err == nil {
						userTTL = val
					}
				}

				if val, exists := os.LookupEnv("URL_TTL"); exists {
					if val, err := strconv.Atoi(val); err == nil {
						urlTTL = val
					}
				}

				instance = &Config{
					UserTTL:   time.Duration(userTTL) * time.Second,
					UserQuota: uint64(userQuota),
					UrlTTL:    time.Duration(urlTTL) * time.Second,
				}
			},
		)
	}
	return instance
}
