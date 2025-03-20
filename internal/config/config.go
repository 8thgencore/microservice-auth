package config

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

// Env represents the environment in which the application is running.
type Env string

const (
	// Local environment.
	Local Env = "local"
	// Dev environment.
	Dev Env = "dev"
	// Prod environment.
	Prod Env = "prod"
)

// Config represents the configuration for the application.
type Config struct {
	Env        Env `env:"ENV" env-default:"local"`
	GRPC       GRPC
	HTTP       HTTPConfig
	JWT        JWTConfig
	TLS        TLSConfig
	Swagger    SwaggerConfig
	Database   DatabaseConfig
	Redis      RedisConfig
	Prometheus PrometheusConfig
	Tracing    TracingConfig
	Admin      AdminConfig
}

// GRPC represents the configuration for the GRPC server.
type GRPC struct {
	Host      string        `env:"GRPC_HOST" env-default:"localhost"`
	Port      int           `env:"GRPC_PORT" env-default:"50051"`
	Transport string        `env:"GRPC_TRANSPORT" env-default:"tcp"`
	Timeout   time.Duration `env:"GRPC_TIMEOUT"`
}

// Address returns the address of the GRPC server in the format "host:port".
func (c *GRPC) Address() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}

// HTTPConfig represents the configuration for the HTTP server.
type HTTPConfig struct {
	Host string `env:"HTTP_HOST" env-default:"0.0.0.0"`
	Port int    `env:"HTTP_PORT" env-default:"8480"`
}

// Address returns the address of the HTTP server in the format "host:port".
func (c *HTTPConfig) Address() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}

// SwaggerConfig represents the configuration for the Swagger server.
type SwaggerConfig struct {
	Host string `env:"HTTP_HOST"    env-default:"0.0.0.0"`
	Port int    `env:"SWAGGER_PORT" env-default:"8490"`
}

// Address returns the address of the Swagger server in the format "host:port".
func (c *SwaggerConfig) Address() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}

// DatabaseConfig represents the configuration for the Postgres database.
type DatabaseConfig struct {
	Host     string `env:"DB_HOST"     env-required:"true"`
	Port     string `env:"DB_PORT"     env-required:"true"`
	User     string `env:"DB_USER"     env-required:"true"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
	Name     string `env:"DB_NAME"     env-required:"true"`
}

// DSN returns the data source name (DSN) for the database.
func (c *DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		c.Host, c.Port, c.Name, c.User, c.Password)
}

// RedisConfig represents the configuration for the Postgres database.
type RedisConfig struct {
	Host              string        `env:"REDIS_HOST"               env-required:"true"`
	Port              int           `env:"REDIS_PORT"               env-required:"true"`
	Password          string        `env:"REDIS_PASSWORD"           env-required:"true"`
	ConnectionTimeout time.Duration `env:"REDIS_CONNECTION_TIMEOUT" env-required:"true"`
	IdleTimeout       time.Duration `env:"REDIS_IDLE_TIMEOUT"       env-required:"true"`
	MaxIdle           int           `env:"REDIS_MAX_IDLE"           env-required:"true"`
}

// Address returns the data source name (Address) for the database.
func (c *RedisConfig) Address() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}

// PrometheusConfig represents the configuration for the Prometheus.
type PrometheusConfig struct {
	Host string `env:"PROMETHEUS_HTTP_HOST" env-default:"0.0.0.0"`
	Port int    `env:"PROMETHEUS_HTTP_PORT" env-default:"9090"`
}

// Address returns the address of the Prometheus server in the format "host:port".
func (c *PrometheusConfig) Address() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}

// TracingConfig represents the configuration for the Jaeger.
type TracingConfig struct {
	Host        string `env:"JAEGER_GRPC_EXPORTER_HOST" env-default:"jaeger"`
	Port        int    `env:"JAEGER_GRPC_EXPORTER_PORT" env-default:"6831"`
	ServiceName string `env:"JAEGER_SERVICE_NAME" env-default:"auth-service"`
}

// Address returns the address of the Tracing server in the format "host:port".
func (c *TracingConfig) Address() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}

// JWTConfig represents the configuration for the JWT.
type JWTConfig struct {
	SecretKey       string        `env:"JWT_SECRET_KEY" env-required:"true"`
	AccessTokenTTL  time.Duration `env:"JWT_ACCESS_TTL" env-default:"15m"`
	RefreshTokenTTL time.Duration `env:"JWT_REFRESH_TTL" env-default:"7d"`
}

// TLSConfig represents the configuration for the TLSConfig.
type TLSConfig struct {
	Enable   bool   `env:"ENABLE_TLS" env-default:"false"`
	CertPath string `env:"TLS_CERT_PATH"`
	KeyPath  string `env:"TLS_KEY_PATH"`
}

// AdminConfig represents the configuration for the admin user.
type AdminConfig struct {
	Email    string `env:"ADMIN_EMAIL"    env-default:"admin@example.com"`
	Password string `env:"ADMIN_PASSWORD" env-default:"admin123"`
	Name     string `env:"ADMIN_NAME"     env-default:"admin"`
}

// NewConfig creates a new instance of Config.
func NewConfig() (*Config, error) {
	configPath := fetchConfigPath()

	cfg := &Config{}
	var err error

	if configPath != "" {
		err = godotenv.Load(configPath)
	} else {
		err = godotenv.Load()
	}
	if err != nil {
		log.Printf("No loading .env file: %v", err)
	}

	if err = cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("error reading env: %w", err)
	}
	log.Printf("Load environment: %s", cfg.Env)

	return cfg, nil
}

func fetchConfigPath() string {
	var configPath string
	flag.StringVar(&configPath, "config", ".env", "Path to config file")

	flag.Parse()

	if configPath == "" {
		configPath = os.Getenv("CONFIG_PATH")
	}

	return configPath
}
