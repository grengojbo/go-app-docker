package controllers

import (
	"fmt"
	"time"
	"net/http"
	"github.com/go-chi/valve"
	"github.com/theplant/appkit/log"
)

// GetSlow
func GetSlow(w http.ResponseWriter, r *http.Request) {

	valve.Lever(r.Context()).Open()
	defer valve.Lever(r.Context()).Close()

	select {
	case <-valve.Lever(r.Context()).Stop():
		logger, ok := log.FromContext(r.Context())
		if ok { logger.Error().Log("msg", "valve is closed. finish up..") }

	case <-time.After(5 * time.Second):
		// The above channel simulates some hard work.
		// We want this handler to complete successfully during a shutdown signal,
		// so consider the work here as some background routine to fetch a long running
		// search query to find as many results as possible, but, instead we cut it short
		// and respond with what we have so far. How a shutdown is handled is entirely
		// up to the developer, as some code blocks are preemptable, and others are not.
		time.Sleep(5 * time.Second)
	}

	w.Write([]byte(fmt.Sprintf("all done.\n")))
}

