package giredis

import (
	"context"
	"strings"

	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/go-redis/redis/v8"
)

type clusterExt func(context.Context, *redis.ClusterClient) error

func NewClusterClient(ctx context.Context, o *Options, exts ...clusterExt) (client *redis.ClusterClient, err error) {

	logger := gilog.FromContext(ctx)

	client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              o.Cluster.Addrs,
		MaxRedirects:       o.Cluster.MaxRedirects,
		ReadOnly:           o.Cluster.ReadOnly,
		RouteByLatency:     o.Cluster.RouteByLatency,
		RouteRandomly:      o.Cluster.RouteRandomly,
		Password:           o.Password,
		MaxRetries:         o.MaxRetries,
		MinRetryBackoff:    o.MinRetryBackoff,
		MaxRetryBackoff:    o.MaxRetryBackoff,
		DialTimeout:        o.DialTimeout,
		ReadTimeout:        o.ReadTimeout,
		WriteTimeout:       o.WriteTimeout,
		PoolSize:           o.PoolSize,
		MinIdleConns:       o.MinIdleConns,
		MaxConnAge:         o.MaxConnAge,
		PoolTimeout:        o.PoolTimeout,
		IdleTimeout:        o.IdleTimeout,
		IdleCheckFrequency: o.IdleCheckFrequency,
	})

	ping := client.Ping(ctx)
	if ping.Err() != nil {
		return nil, ping.Err()
	}

	for _, ext := range exts {
		if err := ext(ctx, client); err != nil {
			panic(err)
		}
	}

	logger.Infof("Connected to Redis Cluster server: %s status: %s", strings.Join(client.Options().Addrs, ","), ping.String())

	return client, err
}

func NewDefaultClusterClient(ctx context.Context, exts ...clusterExt) (*redis.ClusterClient, error) {

	logger := gilog.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewClusterClient(ctx, o, exts...)
}
