package controllers

import (
	"net/http"
	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err            error  `json:"-"`                 // low-level runtime error
	HTTPStatusCode int    `json:"-"`                 // http response status code

	StatusText     string `json:"status"`            // user-level status message
	AppCode        string `json:"code,omitempty"`    // application-specific error code
	ErrorText      string `json:"message,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
return &ErrResponse{
Err:            err,
HTTPStatusCode: 400,
StatusText:     "error",
AppCode: "mess.error.invalid.request",
ErrorText:      err.Error(),
}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "error",
		AppCode: "mess.error.rendering.response",
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "error", AppCode: "mess.error.not.found.resource"}