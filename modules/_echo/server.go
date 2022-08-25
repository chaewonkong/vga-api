package _echo

import (
	"context"
	"log"

	"github.com/labstack/echo"
	"go.uber.org/fx"
)

func NewServer() *echo.Echo {
	e := echo.New()
	return e
}

func RegisterHooks(
	lifecycle fx.Lifecycle,
	e *echo.Echo,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Fatal(e.Start(":8081"))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Fatal("Stop")
			return e.Shutdown(ctx)
		},
	})
}
