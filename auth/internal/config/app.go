package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const (
	prefix            = "Config."
	opNew             = prefix + "New"
	opCheckConfig     = prefix + "CheckConfig"
	opReadEnv         = prefix + "ReadEnv"
	opReadFile        = prefix + "ReadFile"
	opFetchConfigPath = prefix + "FetchConfigPath"

	zeroLifetime = "0s"
)

var (
	ErrConfigPathNotProvided = errors.New("config path didn't provide")
	ErrBadConfigPort         = errors.New("port must be upper than 0")
	ErrBadAddr               = errors.New("bot service's address is required")
	ErrBadSecret             = errors.New("secret is required")
	ErrBadResponseTime       = errors.New("response time must be upper than 0ms")
	ErrBadUserName           = errors.New("username is required")
	ErrBadPassword           = errors.New("password is required")
	ErrBadDBName             = errors.New("db name is required")
	ErrBadDBHost             = errors.New("db host is required")
	ErrBadLifetime           = errors.New("token lifetime is required")
)

type ServerConfig struct {
	Port int    `yaml:"port"`
	Addr string `yaml:"address"`
}

type DBConfig struct {
	User     string `yaml:"user"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
}

type LoggerConfig struct {
	AddSource bool `yaml:"source"`
}

type AuthConfig struct {
	Secret          string        `yaml:"secret"`
	AccessLifetime  time.Duration `yaml:"accessLifetime"`
	RefreshLifetime time.Duration `yaml:"refreshLifetime"`
}

type ClientConfig struct {
	Address    string `yaml:"address"`
	RetryCount int    `yaml:"retryCount"`
}

type Config struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"db"`
	Logger LoggerConfig `yaml:"config"`
	Auth   AuthConfig   `yaml:"auth"`
	Client ClientConfig `yaml:"client"`
}

func New() (cfg *Config, err error) {
	configPath, err := fetchConfigPath()
	if err != nil && !errors.Is(err, ErrConfigPathNotProvided) {
		return nil, fmt.Errorf("%s: %w", opNew, err)
	}

	cfg, err = readFileConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", opNew, err)
	}

	cfg, err = readEnvConfig(cfg)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", opNew, err)
	}

	err = checkConfig(cfg)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", opNew, err)
	}

	return cfg, nil
}

// Валидирует конфиг
func checkConfig(cfg *Config) error {
	if err := checkServerConfig(&cfg.Server); err != nil {
		return fmt.Errorf("%s: %w", opCheckConfig, err)
	}

	if err := checkAuthConfig(&cfg.Auth); err != nil {
		return fmt.Errorf("%s: %w", opCheckConfig, err)
	}

	if err := checkDBConfig(&cfg.DB); err != nil {
		return fmt.Errorf("%s: %w", opCheckConfig, err)
	}

	if err := checkClientConfig(&cfg.Client); err != nil {
		return fmt.Errorf("%s: %w", opCheckConfig, err)
	}

	return nil
}

func checkDBConfig(cfg *DBConfig) error {
	if cfg.Name == "" {
		return ErrBadDBName
	}

	if cfg.Password == "" {
		return ErrBadPassword
	}

	if cfg.User == "" {
		return ErrBadUserName
	}

	return nil
}

func checkServerConfig(cfg *ServerConfig) error {
	if cfg.Port <= 0 {
		return ErrBadConfigPort
	}

	return nil
}

func checkAuthConfig(cfg *AuthConfig) error {
	if cfg.Secret == "" {
		return ErrBadSecret
	}

	if cfg.AccessLifetime.String() == zeroLifetime {
		return ErrBadLifetime
	}

	if cfg.RefreshLifetime.String() == zeroLifetime {
		return ErrBadLifetime
	}

	return nil
}

func checkClientConfig(cfg *ClientConfig) error {
	if cfg.Address == "" {
		return ErrBadAddr
	}

	return nil
}

func readEnvConfig(cfg *Config) (*Config, error) {
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil && cfg.Server.Port == 0 {
		return nil, fmt.Errorf("%s: %w", opReadEnv, err)
	}

	if port != 0 {
		cfg.Server.Port = port
	}

	address := os.Getenv("SERVER_ADDRESS")
	if address != "" {
		cfg.Server.Addr = address
	}

	botAddress := os.Getenv("BOT_ADDRESS")
	if botAddress != "" {
		cfg.Client.Address = botAddress
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword != "" {
		cfg.DB.Password = dbPassword
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser != "" {
		cfg.DB.User = dbUser
	}

	dbName := os.Getenv("DB_NAME")
	if dbName != "" {
		cfg.DB.Name = dbName
	}

	source, err := strconv.ParseBool(os.Getenv("LOGGER_SOURCE"))
	if err != nil {
		source = true
	}

	cfg.Logger.AddSource = source

	return cfg, nil
}

func readFileConfig(configPath string) (*Config, error) {
	var cfg Config
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &cfg, fmt.Errorf("%s: %w", opReadFile, err)
	}

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return &cfg, fmt.Errorf("%s: %w", opReadFile, err)
	}

	return &cfg, nil
}

func fetchConfigPath() (string, error) {
	var path string

	flag.StringVar(&path, "config", "", "path to config file")
	flag.Parse()

	if path == "" {
		err := godotenv.Load()
		if err != nil {
			return "", fmt.Errorf("%s %w", opFetchConfigPath, err)
		}

		path = os.Getenv("CONFIG_PATH")

		if path == "" {
			return "", ErrConfigPathNotProvided
		}
	}

	return path, nil
}

func GetConnStr(user, password, host, name string) string {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", user, password, host, name)
	return connStr
}
