package main

import (
	echo "vga-api/modules/_echo"
	"vga-api/modules/handlers"

	"go.uber.org/fx"
)

func main() {
	modules := fx.Options(
		fx.Provide(echo.NewServer),
		fx.Provide(handlers.NewHandler),
		fx.Invoke(echo.RegisterHooks),
		fx.Invoke(handlers.BindRoutes),
	)
	fx.New(modules).Run()
}
