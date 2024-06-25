package provider

import (
	"github.com/google/wire"

	"gotu-bookstore/internal/app"
	"gotu-bookstore/internal/config"
	"gotu-bookstore/internal/delivery/middleware"
	"gotu-bookstore/internal/repository"
)

var BaseSet = wire.NewSet(
	config.Get,
	LoggerProvider,
	DatabaseProvider,
)

var RepositorySet = wire.NewSet(
	repository.NewUserRepo,
)

var UseCaseSet = wire.NewSet()

var DeliverySet = wire.NewSet(
	middleware.NewMiddleware,
)

var AppSet = wire.NewSet(
	BaseSet,
	RepositorySet,
	UseCaseSet,
	DeliverySet,
	app.AppNew,
)
