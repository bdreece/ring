package ring

import (
	"encoding/json"
	"net/http"
)

type (
	// JSONBinder binds a JSON request body to the provided model type.
    // 
    // See [json.Encoder] for more decoding details.
	JSONBinder[T any] struct {
		// Whether to error on unknown fields
		DisallowUnknownFields bool
	}

	// JSONRenderer renders the provided model type to response body as a JSON
	// string.
    // 
    // See [json.Decoder] for more decoding details.
	JSONRenderer[T any] struct {
		// See [json.Indent]
		Prefix string
		// See [json.Indent]
		Indent string
		// By default, [json.Encoder] escapes certain HTML characters. Set this
		// to true to disable this behavior.
		UnescapeHTML bool
	}
)

// Bind implements Binder.
func (binder *JSONBinder[T]) Bind(r *http.Request, value T) error {
	dec := json.NewDecoder(r.Body)
	if binder.DisallowUnknownFields {
		dec.DisallowUnknownFields()
	}

	return dec.Decode(value)
}

// Render implements Renderer
func (renderer *JSONRenderer[T]) Render(w http.ResponseWriter, value T) error {
	enc := json.NewEncoder(w)
	enc.SetIndent(renderer.Prefix, renderer.Indent)
	enc.SetEscapeHTML(renderer.UnescapeHTML)

	return enc.Encode(value)
}

var (
	_ Binder[any]   = (*JSONBinder[any])(nil)
	_ Renderer[any] = (*JSONRenderer[any])(nil)
)
