// Copyright 2017 Igor Dolzhikov. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"github.com/fatih/color"
	"github.com/grengojbo/go-app-docker/config"
	"github.com/theplant/appkit/log"
	"github.com/theplant/appkit/server"

	"fmt"
	"github.com/grengojbo/go-app-docker/router"
	"github.com/go-chi/valve"
	"github.com/go-chi/chi"
)

var (
	Version     = "0.1.0"
	BuildTime   = "2015-09-20 UTC"
	GitHash     = "c00"
	flagNoColor = flag.Bool("no-color", false, "Disable color output")
)

func main() {
	valv := valve.New()
	baseCtx := valv.Context()

	// Load ENV configuration
	cfg := config.Config
	logger := log.Default()
	if *flagNoColor {
		color.NoColor = true // disables colorized output
	}

	mux := router.Route(logger)



	// Listen and serve handlers
	color.Green("URL: http://%s:%d/", cfg.Host, cfg.Port)
	adr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	server.ListenAndServe(server.Config{Addr: adr}, logger, chi.ServerBaseContext(baseCtx, mux), valv)
}
