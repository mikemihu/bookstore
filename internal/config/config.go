package config

import (
	"os"
	"sync"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"gotu-bookstore/internal/constant"
)

var (
	cfg  *Cfg
	once sync.Once
)

type Cfg struct {
	App     App
	Storage Storage
	Auth    Auth
}

type App struct {
	Address   string
	RateLimit float64
}

type Storage struct {
	Postgres Postgres
	Redis    Redis
}

type Auth struct {
	JwtSecret string
}

type Postgres struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

type Redis struct {
	Address  string
	Password string
}

func Get() *Cfg {
	once.Do(func() {
		// Global timezone
		_ = os.Setenv("TZ", constant.DefaultTimeZone)

		env := os.Getenv("ENV")
		if env == "" {
			env = constant.EnvDevelopment
		}

		// Read-only configuration based on environment
		loadConfig("app." + env)
	})

	return cfg
}

func loadConfig(fileName string) {
	logger := zap.L()

	// Read config use viper
	v := viper.New()
	v.SetConfigName(fileName)
	v.AddConfigPath("./config")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		logger.Fatal("Unable to load config.", zap.Error(err))
	}

	err = v.Unmarshal(&cfg)
	if err != nil {
		logger.Fatal("Unable to unmarshal config.", zap.Error(err))
	}
}
