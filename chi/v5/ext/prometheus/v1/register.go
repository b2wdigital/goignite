package prometheus

import (
	"context"

	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Register(ctx context.Context, instance *chi.Mux) error {

	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("integrating chi with prometheus")

	prometheusRoute := getRoute()

	logger.Infof("configuring prometheus metrics router on %s", prometheusRoute)
	instance.Handle(prometheusRoute, promhttp.Handler())

	instance.Use(promMiddleware)

	logger.Debug("prometheus integrated with echo with success")

	return nil
}
