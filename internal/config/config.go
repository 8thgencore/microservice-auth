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
	Env      Env `env:"ENV" env-default:"local"`
	GRPC     GRPC
	HTTP     HTTPConfig
	TLS      TLS
	Swagger  SwaggerConfig
	Database DatabaseConfig
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
	Host string `env:"SWAGGER_HOST" env-default:"0.0.0.0"`
	Port int    `env:"SWAGGER_PORT" env-default:"8490"`
}

// Address returns the address of the Swagger server in the format "host:port".
func (c *SwaggerConfig) Address() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}

// DatabaseConfig represents the configuration for the Postgres database.
type DatabaseConfig struct {
	Host     string `env:"POSTGRES_HOST"     env-required:"true"`
	Port     string `env:"POSTGRES_PORT"     env-required:"true"`
	User     string `env:"POSTGRES_USER"     env-required:"true"`
	Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
	Name     string `env:"POSTGRES_DB"       env-required:"true"`
}

// DSN returns the data source name (DSN) for the database.
func (c *DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		c.Host, c.Port, c.Name, c.User, c.Password)
}

// TLS represents the configuration for the TLS.
type TLS struct {
	CertPath string `env:"TLS_CERT_PATH"`
	KeyPath  string `env:"TLS_KEY_PATH"`
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
