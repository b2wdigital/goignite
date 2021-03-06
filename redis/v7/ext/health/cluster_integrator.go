package giredishealth

import (
	"context"

	gihealth "github.com/b2wdigital/goignite/v2/health"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/go-redis/redis/v7"
)

type ClusterIntegrator struct {
	options *Options
}

func NewClusterIntegrate(options *Options) *ClusterIntegrator {
	return &ClusterIntegrator{options: options}
}

func NewDefaultClusterIntegrator() *ClusterIntegrator {
	o, err := DefaultOptions()
	if err != nil {
		gilog.Fatalf(err.Error())
	}

	return NewClusterIntegrate(o)
}

func (i *ClusterIntegrator) Register(ctx context.Context, client *redis.ClusterClient) error {

	logger := gilog.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating redis in health")

	checker := NewClusterClientChecker(client)
	hc := gihealth.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	gihealth.Add(hc)

	logger.Debug("redis successfully integrated in health")

	return nil
}
