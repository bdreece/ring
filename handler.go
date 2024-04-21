package ring

import "net/http"

// Handler composes a Binder, Controller, and Renderer to implement a
// [http.Handler]
type Handler[Req, Res any] struct {
	controller   Controller[Req, Res]
	binder       Binder[Req]
	renderer     Renderer[Res]
	errorHandler ErrorHandler
}

// HandlerOptions provide the required dependencies for a Handler
type HandlerOptions[Req, Res any] struct {
	Controller   Controller[Req, Res]
	Binder       Binder[Req]
	Renderer     Renderer[Res]
	ErrorHandler ErrorHandler
}

// ServeHTTP implements [http.Handler]
func (handler *Handler[Req, Res]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	model := *new(Req)
	if err := handler.binder.Bind(r, model); err != nil {
        handler.errorHandler.Handle(w, err)
        return
	}

	res, err := handler.controller.Execute(r.Context(), model)
	if err != nil {
        handler.errorHandler.Handle(w, err)
        return
	}

	if err = handler.renderer.Render(w, res); err != nil {
        handler.errorHandler.Handle(w, err)
        return
	}
}

// NewHandler creates a Handler with the provided opts
func NewHandler[Req, Res any](opts *HandlerOptions[Req, Res]) *Handler[Req, Res] {
	return &Handler[Req, Res]{
		controller:   opts.Controller,
		binder:       opts.Binder,
		renderer:     opts.Renderer,
		errorHandler: opts.ErrorHandler,
	}
}
