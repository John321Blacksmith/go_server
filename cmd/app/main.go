// this package initializes
// an entrypoint

package main

import (
	"fmt"
	"media_api/config"
	"media_api/internal/app"

	"golang.org/x/exp/slog"
)

// import the configs and
// apply them to the servers launch
func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		slog.Info(fmt.Sprint(err))
	}
	err = app.Run(cfg)
	if err != nil {
		slog.Info(fmt.Sprintf("During application startup, an error occurred: %v", err))
	}
}
