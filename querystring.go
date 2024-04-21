package ring

import (
	"net/http"

	"github.com/pasztorpisti/qs"
)

// QueryStringBinder binds the request URL query string to the provided model
// using the \`qs:""\` struct tag.
//
// See [qs.QSUnmarshaler] for more decoding details.
type QueryStringBinder[T any] struct {
    Options *qs.UnmarshalOptions
}

// Bind implements Binder.
func (binder *QueryStringBinder[T]) Bind(r *http.Request, model T) error {
    return qs.NewUnmarshaler(binder.Options).Unmarshal(model, r.URL.RawQuery)
}

var _ Binder[any] = (*QueryStringBinder[any])(nil)
