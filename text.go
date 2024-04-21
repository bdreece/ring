package ring

import (
	"io/fs"
	"net/http"
	"text/template"
)

// TextRenderer renders the provided model to the response body as text using
// the named template.
//
// See [template.Template] for more rendering details.
type TextRenderer[T any] struct {
    // The template collection
    *template.Template

	// The name of the target template
	Name string
	// Set to true to re-parse the target template before rendering. This is
	// useful when nesting template layouts.
	Reparse bool
	// Source must be included when reparsing targets, otherwith the renderer
	// will panic
	Source fs.FS
}

// Render implements Renderer.
func (renderer TextRenderer[T]) Render(w http.ResponseWriter, value T) error {
    t := template.Must(renderer.Clone())
	if renderer.Reparse {
		if renderer.Source == nil {
			panic("Source cannot be nil when reparsing target template")
		}
		t = template.Must(t.ParseFS(renderer.Source, renderer.Name))
	}

	return t.ExecuteTemplate(w, renderer.Name, value)
}

var _ Renderer[any] = (*TextRenderer[any])(nil)
