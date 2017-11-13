package router

import (
	"net/http"
	"github.com/theplant/appkit/server"
	"github.com/theplant/appkit/log"
	"github.com/theplant/appkit/contexts"
	"github.com/go-chi/chi"
	//"github.com/theplant/appkit/contexts/trace"
	"github.com/grengojbo/go-app-docker/controllers"
	"fmt"
	"github.com/go-chi/render"
	"github.com/go-chi/docgen"
	"flag"
)

var flagRoutes = flag.Bool("routes", false, "Generate router documentation")

func Route(logger log.Logger) (http.Handler) {
	r := chi.NewRouter()

	r.Use(AppMiddleware(logger))

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-type", "text/plain; charset=utf-8")
		w.Write([]byte("ok"))
	})

	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/slow", controllers.GetSlow)

	if *flagRoutes {
		// fmt.Println(docgen.JSONRoutesDoc(r))
		fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
			ProjectPath: "github.com/go-chi/chi",
			Intro:       "Welcome to the chi/_examples/rest generated docs.",
		}))
		//return
	}

	return r
}

func AppMiddleware(logger log.Logger) func(http.Handler) http.Handler {
	return server.Compose(
		server.ETag,
		server.Recovery,
		server.LogRequest,
		log.WithLogger(logger),
		//trace.WithRequestTrace,
		contexts.WithHTTPStatus,
	)
}

// This is entirely optional, but I wanted to demonstrate how you could easily
// add your own logic to the render.Respond method.
func init() {
	render.Respond = func(w http.ResponseWriter, r *http.Request, v interface{}) {
		if err, ok := v.(error); ok {

			// We set a default error status response code if one hasn't been set.
			if _, ok := r.Context().Value(render.StatusCtxKey).(int); !ok {
				w.WriteHeader(400)
			}

			// We log the error
			fmt.Printf("Logging err: %s\n", err.Error())

			// We change the response to not reveal the actual error message,
			// instead we can transform the message something more friendly or mapped
			// to some code / language, etc.
			render.DefaultResponder(w, r, render.M{"status": "error"})
			return
		}

		render.DefaultResponder(w, r, v)
	}
}