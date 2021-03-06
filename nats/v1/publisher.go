package ginats

import (
	"context"

	"github.com/nats-io/nats.go"
)

type PublisherMiddleware interface {
	Before(context.Context, *nats.Conn, *nats.Msg) (context.Context, error)
	After(context.Context) error
}

type Publisher struct {
	conn        *nats.Conn
	options     *Options
	middlewares []PublisherMiddleware
}

func NewPublisher(ctx context.Context, options *Options, middlewares ...PublisherMiddleware) (*Publisher, error) {
	conn, err := NewConnection(ctx, options)
	if err != nil {
		return nil, err
	}
	return &Publisher{conn, options, middlewares}, nil
}

func NewDefaultPublisher(ctx context.Context, middlewares ...PublisherMiddleware) (*Publisher, error) {
	options, err := DefaultOptions()
	if err != nil {
		return nil, err
	}
	return NewPublisher(ctx, options, middlewares...)
}

func (p *Publisher) Publish(ctx context.Context, msg *nats.Msg) error {

	for _, middleware := range p.middlewares {
		ctxx, err := middleware.Before(ctx, p.conn, msg)
		if err != nil {
			return err
		}
		defer func() {
			middleware.After(ctxx)
		}()
	}

	return p.conn.PublishMsg(msg)
}

func (p *Publisher) Conn() *nats.Conn {
	return p.conn
}
