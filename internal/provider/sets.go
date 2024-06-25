package provider

import (
	"gotu-bookstore/internal/app"
	"gotu-bookstore/internal/config"
	"gotu-bookstore/internal/delivery"
	"gotu-bookstore/internal/delivery/middleware"
	"gotu-bookstore/internal/repository"
	"gotu-bookstore/internal/usecase"

	"github.com/google/wire"
)

var BaseSet = wire.NewSet(
	config.Get,
	LoggerProvider,
	DatabaseProvider,
	AuthJWTProvider,
)

var RepositorySet = wire.NewSet(
	repository.NewUserRepo,
)

var UseCaseSet = wire.NewSet(
	usecase.NewUserUC,
)

var DeliverySet = wire.NewSet(
	middleware.NewMiddleware,
	delivery.NewAuthDelivery,
	delivery.NewUserDelivery,
)

var AppSet = wire.NewSet(
	BaseSet,
	RepositorySet,
	UseCaseSet,
	DeliverySet,
	app.AppNew,
)
