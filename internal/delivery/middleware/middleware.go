package middleware

import (
	"gotu-bookstore/internal/config"

	"go.uber.org/zap"
)

type Middleware struct {
	cfg    *config.Cfg
	logger *zap.Logger
}

func NewMiddleware(
	cfg *config.Cfg,
	logger *zap.Logger,
) *Middleware {
	return &Middleware{
		cfg:    cfg,
		logger: logger,
	}
}
