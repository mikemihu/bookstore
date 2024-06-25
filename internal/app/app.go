package app

import (
	"context"
	"errors"
	"gotu-bookstore/internal"
	"net/http"
	"time"

	"gotu-bookstore/internal/config"
	"gotu-bookstore/internal/delivery/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type App struct {
	cfg          *config.Cfg
	logger       *zap.Logger
	middleware   *middleware.Middleware
	gin          *gin.Engine
	server       *http.Server
	authDelivery internal.AuthDelivery
	userDelivery internal.UserDelivery
	bookDelivery internal.BookDelivery
}

func AppNew(
	cfg *config.Cfg,
	logger *zap.Logger,
	middleware *middleware.Middleware,
	authDelivery internal.AuthDelivery,
	userDelivery internal.UserDelivery,
	bookDelivery internal.BookDelivery,
) *App {
	return &App{
		cfg:          cfg,
		logger:       logger,
		gin:          gin.Default(),
		middleware:   middleware,
		authDelivery: authDelivery,
		userDelivery: userDelivery,
		bookDelivery: bookDelivery,
	}
}

const (
	_ShutdownTimeout = 5 * time.Second
)

func (a *App) Run() error {
	a.logger.Info("starting app")

	// routing
	a.registerRoutes()

	a.logger.Info("app listening", zap.String("address", a.cfg.App.Address))

	a.server = &http.Server{
		Addr:    a.cfg.App.Address,
		Handler: a.gin.Handler(),
	}
	// start server
	err := a.server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		// ignore the shutdown call error
		return nil
	}
	if err != nil {
		return err
	}

	return nil
}

// Stop gracefully stop the app
func (a *App) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), _ShutdownTimeout)
	defer cancel()

	err := a.server.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}
