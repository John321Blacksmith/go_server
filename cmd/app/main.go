// this package initializes
// an entrypoint

package main

import (
	"log"

	"media_api/config"
	"media_api/internal/app"
)

// import the configs and
// apply them to the servers launch
func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatal(err)
	}
	app.Run(cfg)
}
