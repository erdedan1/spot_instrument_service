package main

import (
	"spot_instrument_service/config"
	"spot_instrument_service/internal/app"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	orderService := app.New(cfg)
	if err := orderService.Run(); err != nil {
		panic(err)
	}
}
