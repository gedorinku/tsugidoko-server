package worker

import (
	"context"
	"time"

	"github.com/gedorinku/tsugidoko-server/app/config"
	"github.com/gedorinku/tsugidoko-server/app/di"
	"google.golang.org/grpc/grpclog"
)

var stop = make(chan struct{})

const interval = 30 * time.Second

// Run runs worker
func Run(cfg *config.Config, store di.StoreComponent) {
	lifeTime := time.Duration(cfg.UserPositionLifeSecond) * time.Second

	go func() {
		defer func() {
			if err := recover(); err != nil {
				grpclog.Error(err)
				Run(cfg, store)
			}
			grpclog.Infoln("worker stopped\n")
		}()

		for {
			select {
			case <-time.Tick(interval):
				err := expireUserPosition(store, lifeTime)
				if err != nil {
					grpclog.Error(err)
				}
			case <-stop:
				return
			}
		}
	}()
}

// Close stops worker
func Close() {
	stop <- struct{}{}
}

func expireUserPosition(store di.StoreComponent, lifeTime time.Duration) error {
	_, err := store.UserPositionStore(context.Background()).Expire(lifeTime)
	if err != nil {
		return err
	}
	return nil
}
