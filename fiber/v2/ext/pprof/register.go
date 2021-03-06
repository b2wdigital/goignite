package gifiberpprof

import (
	"context"

	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func Register(ctx context.Context, app *fiber.App) error {
	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("enabling pprof middleware in fiber")

	app.Use(pprof.New())

	logger.Debug("pprof middleware successfully enabled in fiber")

	return nil
}
