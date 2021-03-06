package gigocloud

import (
	"context"
	"strings"

	gilog "github.com/b2wdigital/goignite/v2/log"
	"gocloud.dev/pubsub"
	_ "gocloud.dev/pubsub/awssnssqs"
	_ "gocloud.dev/pubsub/gcppubsub"
	_ "gocloud.dev/pubsub/kafkapubsub"
	_ "gocloud.dev/pubsub/mempubsub"
)

// Queue types
const (
	SQS    = "sqs"
	SNS    = "sns"
	KAFKA  = "kafka"
	NATS   = "nats"
	PUBSUB = "pubsub"
)

// NewTopic start a new topic for send message
func NewTopic(ctx context.Context, o *Options) (*pubsub.Topic, error) {

	logger := gilog.FromContext(ctx)

	addResource(o)

	if o.Region != "" {
		o.Resource = appendRegion(o.Resource, o.Region)
	}

	topic, err := pubsub.OpenTopic(ctx, o.Resource)
	if err != nil {
		return nil, err
	}

	logger.Infof("open topic for send message")
	return topic, nil

}

// NewDefaultTopic ..
func NewDefaultTopic(ctx context.Context) (*pubsub.Topic, error) {

	logger := gilog.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewTopic(ctx, o)
}

// NewSubscription ..
func NewSubscription(ctx context.Context, o *Options) (*pubsub.Subscription, error) {

	logger := gilog.FromContext(ctx)

	addResource(o)

	if o.Region != "" {
		o.Resource = appendRegion(o.Resource, o.Region)
	}

	subscription, err := pubsub.OpenSubscription(ctx, o.Resource)
	if err != nil {
		return nil, err
	}

	logger.Infof("open subscription for listen")
	return subscription, nil

}

// NewDefaultSubscription ..
func NewDefaultSubscription(ctx context.Context) (*pubsub.Subscription, error) {

	logger := gilog.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewSubscription(ctx, o)
}

func addResource(o *Options) {
	switch strings.ToLower(o.Type) {
	case SQS:
		o.Resource = addSQSResource(o.Resource)
	case SNS:
		o.Resource = addSNSResource(o.Resource)
	case KAFKA:
		// TODO: https://gocloud.dev/howto/pubsub/publish/#kafka
	case NATS:
		// TODO: https://gocloud.dev/howto/pubsub/publish/#nats
	case PUBSUB:
		// TODO: https://gocloud.dev/howto/pubsub/publish/#gcp
	default:
		o.Resource = addMEMResource(o.Resource)
	}
}

func addMEMResource(topicName string) string {
	return "mem://" + topicName
}

func addSQSResource(url string) string {
	newURL := strings.Replace(url, "https://", "awssqs://", -1)
	return newURL
}

func addSNSResource(arn string) string {
	return "awssns:///" + arn
}

func appendRegion(add, region string) string {
	return add + "?region=" + region
}
