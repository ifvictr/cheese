package cheese

import (
	"os"
	"strconv"
)

type Config struct {
	BotToken      string
	Port          int
	RedisURL      string
	SigningSecret string
}

func NewConfig() *Config {
	return &Config{
		BotToken:      getEnv("SLACK_CLIENT_BOT_TOKEN", ""),
		Port:          getEnvAsInt("PORT", 3000),
		RedisURL:      getEnv("REDIS_URL", "redis://localhost:6379/0"),
		SigningSecret: getEnv("SLACK_CLIENT_SIGNING_SECRET", ""),
	}
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}

	return defaultValue
}

func getEnvAsInt(name string, defaultValue int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultValue
}
