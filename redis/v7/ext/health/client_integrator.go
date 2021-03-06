package giredishealth

import (
	"context"

	gihealth "github.com/b2wdigital/goignite/v2/health"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/go-redis/redis/v7"
)

type ClientIntegrator struct {
	options *Options
}

func NewClientIntegrator(options *Options) *ClientIntegrator {
	return &ClientIntegrator{options: options}
}

func NewDefaultClientIntegrator() *ClientIntegrator {
	o, err := DefaultOptions()
	if err != nil {
		gilog.Fatalf(err.Error())
	}

	return NewClientIntegrator(o)
}

func (i *ClientIntegrator) Register(ctx context.Context, client *redis.Client) error {

	logger := gilog.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating redis in health")

	checker := NewClientChecker(client)
	hc := gihealth.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	gihealth.Add(hc)

	logger.Debug("redis successfully integrated in health")

	return nil
}
