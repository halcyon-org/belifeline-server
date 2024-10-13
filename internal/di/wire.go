//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/halcyon-org/kizuna/internal/adapter/api"
	"github.com/halcyon-org/kizuna/internal/adapter/controller"
	"github.com/halcyon-org/kizuna/internal/adapter/interceptor"
	"github.com/halcyon-org/kizuna/internal/adapter/repository/config"
	repoent "github.com/halcyon-org/kizuna/internal/adapter/repository/ent"
	infraent "github.com/halcyon-org/kizuna/internal/infrastructure/ent"
	"github.com/halcyon-org/kizuna/internal/usecase"
)

// Adapter
var repositorySet = wire.NewSet(
	config.NewConfigRepository,
	repoent.NewAdminUserRepository,
	repoent.NewClientInformationRepository,
	repoent.NewKoyoInformationRepository,
	repoent.NewExternalInformationRepository,
)

var adapterSet = wire.NewSet(
	api.NewAdminServiceHandler,
	api.NewProviderServiceHandler,
	api.NewExternalInformationServiceHandler,
	api.NewKoyoServiceHandler,
	api.NewServerServiceHandler,
	interceptor.NewAuthInterceptorAdapter,
)

var controllerSet = wire.NewSet(
	controller.NewBeLifelineController,
)

// Infrastructure
var infrastructureSet = wire.NewSet(
	infraent.InitDB,
)

// Usecase
var usecaseSet = wire.NewSet(
	usecase.NewAuthUsecase,
	usecase.NewKoyoInformationUsecase,
	usecase.NewClientInformationUsecase,
)

type ControllersSet struct {
	BeLifelineController controller.BeLifelineController
}

func InitializeControllerSet() (*ControllersSet, error) {
	wire.Build(
		repositorySet,
		adapterSet,
		controllerSet,
		infrastructureSet,
		usecaseSet,
		wire.Struct(new(ControllersSet), "*"),
	)
	return nil, nil
}
