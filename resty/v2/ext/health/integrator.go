package girestyhealth

import (
	"context"

	gihealth "github.com/b2wdigital/goignite/v2/health"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/go-resty/resty/v2"
)

type Integrator struct {
	options *Options
}

func NewIntegrator(options *Options) *Integrator {
	return &Integrator{options: options}
}

func NewDefaultIntegrator() *Integrator {
	o, err := DefaultOptions()
	if err != nil {
		gilog.Fatalf(err.Error())
	}

	return NewIntegrator(o)
}

func (i *Integrator) Register(ctx context.Context, client *resty.Client) error {

	logger := gilog.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating resty in health")

	checker := NewChecker(client, i.options)
	hc := gihealth.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	gihealth.Add(hc)

	logger.Debug("resty successfully integrated in health")

	return nil
}
