package ring

import "net/http"

// A Binder binds the [http.Request] to the given model
type Binder[T any] interface {
    Bind(r *http.Request, model T) error
}
