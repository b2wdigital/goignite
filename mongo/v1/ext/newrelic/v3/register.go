package gimongonewrelic

import (
	"context"

	gilog "github.com/b2wdigital/goignite/v2/log"
	gimongo "github.com/b2wdigital/goignite/v2/mongo/v1"
	"github.com/newrelic/go-agent/v3/integrations/nrmongo"
)

func Register(ctx context.Context, conn *gimongo.Conn) error {

	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("integrating mongo in newrelic")

	nrMon := nrmongo.NewCommandMonitor(nil)

	conn.ClientOptions.SetMonitor(nrMon)

	logger.Debug("mongo successfully integrated in newrelic")

	return nil
}
