package app

import (
	"context"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/gedorinku/tsugidoko-server/app/config"
	"github.com/gedorinku/tsugidoko-server/app/di"
	"github.com/gedorinku/tsugidoko-server/app/interceptor"
	"github.com/gedorinku/tsugidoko-server/app/server"
)

// Run starts the grapiserver.
func Run() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	boil.DebugMode = cfg.DebugLog

	store, err := di.NewStoreComponent(cfg)
	if err != nil {
		return err
	}

	err = store.UserPositionStore(context.Background()).ResetUserPosition()
	if err != nil {
		return err
	}

	authorizator := interceptor.NewAuthorizator(store)

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithGrpcServerUnaryInterceptors(
			authorizator.UnaryServerInterceptor(),
		),
		grapiserver.WithServers(
			server.NewUserServiceServer(store),
			server.NewSessionServiceServer(store),
			server.NewClassRoomServiceServer(store),
			server.NewUserPositionServiceServer(store),
		),
		grapiserver.WithGrpcAddr("tcp", ":4000"),
	)
	return s.Serve()
}
