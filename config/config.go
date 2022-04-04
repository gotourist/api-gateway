package config

import (
	"github.com/joho/godotenv"
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	HttpPort string

	CtxTimeout int

	PostServiceHost string
	PostServicePort int

	CollectServiceHost string
	CollectServicePort int
}

// Load ...
func Load() Config {
	_ = godotenv.Load()

	config := Config{}

	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8050"))

	config.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	config.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "0.0.0.0"))
	config.PostServicePort = cast.ToInt(getOrReturnDefault("POST_SERVICE_PORT", 9090))

	config.CollectServiceHost = cast.ToString(getOrReturnDefault("COLLECT_SERVICE_HOST", "0.0.0.0"))
	config.CollectServicePort = cast.ToInt(getOrReturnDefault("COLLECT_SERVICE_PORT", 9090))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
