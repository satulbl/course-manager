package config

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App      App
	HTTP     HTTP
	Postgres Postgres
	Redis    Redis
	JWT      JWT
	Cleanup  Cleanup
}

type App struct {
	Env      string     `envconfig:"APP_ENV" default:"local"`
	LogLevel slog.Level `envconfig:"LOG_LEVEL" default:"info"`
	LogFile  string     `envconfig:"APP_LOG" default:"app.log"`
}

type HTTP struct {
	Port            string        `envconfig:"HTTP_PORT" default:"8080"`
	ReadTimeout     time.Duration `envconfig:"HTTP_READ_TIMEOUT" default:"30s"`
	WriteTimeout    time.Duration `envconfig:"HTTP_WRITE_TIMEOUT" default:"30s"`
	ShutdownTimeout time.Duration `envconfig:"HTTP_SHUTDOWN_TIMEOUT" default:"10s"`
}

type Postgres struct {
	Host     string `envconfig:"POSTGRES_HOST" default:"localhost"`
	Port     string `envconfig:"POSTGRES_PORT"`
	User     string `envconfig:"POSTGRES_USER"`
	Password string `envconfig:"POSTGRES_PASSWORD"`
	DataBase string `envconfig:"POSTGRES_DB"`
	SSLMode  string `envconfig:"POSTGRES_SSLMODE"`
}

type Redis struct {
	Host     string `envconfig:"REDIS_HOST" default:"localhost"`
	Port     string `envconfig:"REDIS_PORT" default:"6379"`
	Password string `envconfig:"REDIS_PASSWORD"`
	DB       int    `envconfig:"REDIS_DB"`
}

type JWT struct {
	Secret     string `envconfig:"JWT_SECRET" required:"true"`
	Issuer     string `envconfig:"JWT_ISSUER" default:"courses-manager-project"`
	AccessTTL  time.Duration `envconfig:"JWT_ACCESS_TTL" default:"15m"`
	RefreshTTL time.Duration `envconfig:"JWT_REFRESH_TTL" default:"72h"`
}

type Cleanup struct {
	Enabled       bool          `envconfig:"CLEANUP_ENABLED" default:"true"`
	Interval      time.Duration `envconfig:"CLEANUP_INTERVAL" default:"30m"`
	UnattachedTTL time.Duration `envconfig:"CLEANUP_UNATTACHED_TTL" default:"24h"`
}

func (c *Config) validate() error {
	if c.Cleanup.Interval <= 0 {
		return fmt.Errorf("CLEANUP_INTERVAL must be positive")
	}
	if c.Cleanup.UnattachedTTL <= 0 {
		return fmt.Errorf("CLEANUP_UNATTACHED_TTL must be positive")
	}

	return nil
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("load .env: %w", err)
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}
