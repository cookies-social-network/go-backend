package initiator

import (
	"csn-backend/internal/config"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type App struct {
	cfg    *config.Config
	logger *zap.Logger
	web    *fiber.App
}

func NewApp(cfg *config.Config, logger *zap.Logger) (*App, error) {
	web := fiber.New()

	web.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	return &App{
		logger: logger,
		cfg:    cfg,
		web:    web,
	}, nil
}

func (app *App) Run() {
	errListen := app.web.Listen(":3000")
	if errListen != nil {
		app.logger.Fatal("failed to start listening to the port", zap.Error(errListen))
		return
	}
}
