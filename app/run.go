package app

import (
	"github.com/gedorinku/tsugidoko-server/app/config"
	"github.com/gedorinku/tsugidoko-server/app/di"
	"github.com/gedorinku/tsugidoko-server/app/server"
	"github.com/izumin5210/grapi/pkg/grapiserver"
)

// Run starts the grapiserver.
func Run() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	store, err = di.NewStoreComponent(cfg)
	if err != nil {
		return err
	}

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewUserServiceServer(store),
		),
	)
	return s.Serve()
}
