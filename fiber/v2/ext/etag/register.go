package gifiberetag

import (
	"context"

	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
)

func Register(ctx context.Context, app *fiber.App) error {
	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("enabling etag middleware in fiber")

	app.Use(etag.New())

	logger.Debug("etag middleware successfully enabled in fiber")

	return nil
}
