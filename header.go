package ring

import (
	"maps"
	"net/http"

	"github.com/mozillazg/go-httpheader"
)

type (
    // HeaderBinder binds the request headers to the provided model using the
    // \`header:""\` struct tag.
    //
    // See [httpheader.Decode] for more decoding details.
    HeaderBinder[T any] struct {}

    // HeaderRenderer renders the provided model to the response headers using
    // the \`header:""\` struct tag. Headers are merged into the pre-existing
    // header map present in the [http.ResponseWriter].
    // 
    // See [httpheader.Encode] for more encoding details.
    HeaderRenderer[T any] struct {}
)

// Bind implements Binder.
func (HeaderBinder[T]) Bind(r *http.Request, value T) error {
    return httpheader.Decode(r.Header, value)
}

// Render implements Renderer.
func (HeaderRenderer[T]) Render(w http.ResponseWriter, value T) error {
    headers, err := httpheader.Encode(value)
    if err != nil {
        return err
    }

    maps.Copy(w.Header(), headers)
    return nil
}

var (
    _ Binder[any] = (*HeaderBinder[any])(nil)
    _ Renderer[any] = (*HeaderRenderer[any])(nil)
)
