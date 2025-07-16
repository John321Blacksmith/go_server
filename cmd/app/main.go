// this package initializes
// an entrypoint

package main

import (
	cfg "media_api/config"
	app "media_api/internal/app"

	"golang.org/x/exp/slog"
)

func main() {
	cfg, err := cfg.NewConfig()
	if err != nil {
		slog.Info("cannot fetch configs: %v", err)
	}
	err = app.Run(cfg)
	if err != nil {
		slog.Error("error occurred setting up HTTP server: %v", err)
	}
}
