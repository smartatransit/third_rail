package main

import (
	"github.com/smartatransit/third_rail/pkg/api"
)

func main() {
	app := &api.App{}
	app.MountAndServe()
}
