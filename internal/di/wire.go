//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/halcyon-org/kizuna/internal/adapter/api"
	"github.com/halcyon-org/kizuna/internal/adapter/controller"
	"github.com/halcyon-org/kizuna/internal/adapter/repository/config"
)

// Adapter
var repositorySet = wire.NewSet(
	config.NewConfigRepository,
)

var adapterSet = wire.NewSet(
	api.NewBeLifelineServiceHandler,
)

var controllerSet = wire.NewSet(
	controller.NewBeLifelineController,
)

// Infrastructure
// var infrastructureSet = wire.NewSet()

// Usecase
// var usecaseSet = wire.NewSet()

type ControllersSet struct {
	BeLifelineController controller.BeLifelineController
}

func InitializeControllerSet() (*ControllersSet, error) {
	wire.Build(
		repositorySet,
		adapterSet,
		controllerSet,
		wire.Struct(new(ControllersSet), "*"),
	)
	return nil, nil
}
