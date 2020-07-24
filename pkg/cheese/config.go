package cheese

import (
	"os"
)

type Config struct {
	BotToken string
	RedisURL string
}

func NewConfig() *Config {
	return &Config{
		BotToken: getEnv("SLACK_CLIENT_BOT_TOKEN", ""),
		RedisURL: getEnv("REDIS_URL", "redis://localhost:6379/0"),
	}
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}

	return defaultValue
}
