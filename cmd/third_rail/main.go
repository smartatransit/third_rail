package main

import (
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	"github.com/smartatransit/third_rail/pkg/api"
)

func main() {
	var options api.Options
	_, err := flags.Parse(&options)

	if err != nil {
		log.Fatal(err)
	}

	app := &api.App{Options: options}
	if err = app.Initialize(); err != nil {
		log.Fatal(err)
	}
	app.Start(nil)
}
