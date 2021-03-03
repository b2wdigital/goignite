package recover

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Register(ctx context.Context, app *fiber.App) error {
	if IsEnabled() {
		app.Use(recover.New())
	}

	return nil
}
