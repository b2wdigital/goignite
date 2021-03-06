package gifiberrequestid

import (
	"context"

	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Register(ctx context.Context, app *fiber.App) error {
	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("enabling requestID middleware in fiber")

	app.Use(requestid.New())

	logger.Debug("requestID middleware successfully enabled in fiber")

	return nil
}
