package ring

import "net/http"

// A Renderer renders the given model to the response body
type Renderer[T any] interface {
    Render(w http.ResponseWriter, model T) error
}
