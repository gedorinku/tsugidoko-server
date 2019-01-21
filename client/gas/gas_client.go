package gas

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"github.com/gedorinku/tsugidoko-server/client"
)

type gasClientImpl struct {
	ctx context.Context
	url *url.URL
}

// NewGASClient creates new gas client
func NewGASClient(ctx context.Context, url *url.URL) client.GASClient {
	return &gasClientImpl{
		ctx: ctx,
		url: url,
	}
}

func (c *gasClientImpl) GetBuildings() (*client.GetBuildingsResponse, error) {
	resp, err := http.Get(c.url.String())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()

	res, err := c.decodeResponse(resp)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *gasClientImpl) decodeResponse(resp *http.Response) (*client.GetBuildingsResponse, error) {
	dec := json.NewDecoder(resp.Body)
	res := &client.GetBuildingsResponse{}
	err := dec.Decode(res)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return res, nil
}
