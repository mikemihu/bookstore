//go:build wireinject
// +build wireinject

package provider

import (
	"gotu-bookstore/internal/app"

	"github.com/google/wire"
)

func ProvideApp() *app.App {
	wire.Build(AppSet)

	// Return any struct that exist inside the build
	return &app.App{}
}
