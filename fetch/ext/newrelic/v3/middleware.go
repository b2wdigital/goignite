package gifetchnewrelic

import (
	"context"
	"net/http"

	gifetch "github.com/b2wdigital/goignite/v2/fetch"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type middleware struct {
}

func (m *middleware) OnBeforeRequest(ctx context.Context, o gifetch.Options) context.Context {
	reqHTTP, _ := http.NewRequest(o.Method, o.Url, nil)
	txn := newrelic.FromContext(ctx)
	s := newrelic.StartExternalSegment(txn, reqHTTP)
	ctx = context.WithValue(ctx, "s", s)
	return ctx
}

func (m *middleware) OnAfterRequest(ctx context.Context, o gifetch.Options, r gifetch.Response) {
	s := ctx.Value("s").(*newrelic.ExternalSegment)
	s.End()
}

func New() gifetch.Middleware {
	return &middleware{}
}