package di

import (
	"context"
	"net/url"

	"github.com/pkg/errors"

	"github.com/gedorinku/tsugidoko-server/app/config"
	"github.com/gedorinku/tsugidoko-server/client"
	"github.com/gedorinku/tsugidoko-server/client/gas"
)

// ClientComponent is an interface of clients
type ClientComponent interface {
	GASClient(ctx context.Context) client.GASClient
}

// NewClientComponent creates new client component
func NewClientComponent(cfg *config.Config) (ClientComponent, error) {
	u, err := url.Parse(cfg.GASURL)
	if err != nil {
		return nil, errors.Wrap(err, "cannot parse gas url")
	}

	return &clientComponentImpl{
		url: u,
	}, nil
}

type clientComponentImpl struct {
	url *url.URL
}

func (c *clientComponentImpl) GASClient(ctx context.Context) client.GASClient {
	return gas.NewGASClient(ctx, c.url)
}
