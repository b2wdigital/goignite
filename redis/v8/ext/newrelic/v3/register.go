package newrelic

import (
	"context"

	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/go-redis/redis/v8"
)

func Register(ctx context.Context, client *redis.Client) error {

	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("integrating redis with newrelic")

	client.AddHook(NewHook(client.Options()))
	logger.Debug("redis integrated with newrelic with success")

	return nil
}
