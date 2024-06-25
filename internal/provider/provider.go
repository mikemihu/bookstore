package provider

import (
	"fmt"
	"gotu-bookstore/pkg/authentication"
	"sync"

	"gotu-bookstore/internal/config"
	"gotu-bookstore/internal/constant"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	logger     *zap.Logger
	loggerOnce sync.Once

	db     *gorm.DB
	dbOnce sync.Once
)

func LoggerProvider() *zap.Logger {
	loggerOnce.Do(func() {
		logger, _ = zap.NewDevelopment()
		_ = zap.ReplaceGlobals(logger)
	})
	return logger
}

func DatabaseProvider(cfg *config.Cfg) *gorm.DB {
	dbOnce.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
			cfg.Storage.Postgres.Host,
			cfg.Storage.Postgres.Port,
			cfg.Storage.Postgres.User,
			cfg.Storage.Postgres.Password,
			cfg.Storage.Postgres.Database,
			constant.DefaultTimeZone,
		)

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			TranslateError: true,
		})
		if err != nil {
			logger.Fatal("failed open database", zap.Error(err))
		}

		db = db.Debug()
	})
	return db
}

func AuthJWTProvider(cfg *config.Cfg) authentication.AuthJWT {
	return authentication.NewAuthJWT([]byte(cfg.Auth.JwtSecret))
}
