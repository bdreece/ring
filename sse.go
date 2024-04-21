package ring

import (
	"fmt"
	"net/http"
)

// SSERenderer renders a channel of model data to the response body as a
// server-sent event stream.
type SSERenderer[T any] struct {
    // Serializes the model to a string
    TextMarshaler func(T) ([]byte, error)
}

// Render implements Renderer.
func (renderer *SSERenderer[T]) Render(w http.ResponseWriter, ch <-chan T) error {
    for {
        model, ok := <-ch
        if !ok {
            return nil
        }

        text, err := renderer.TextMarshaler(model)
        if err != nil {
            return err
        }

        msg := fmt.Sprintf("data: %s", text)
        if _, err = w.Write([]byte(msg)); err != nil {
            return err
        }
    }
}

var _ Renderer[<-chan any] = (*SSERenderer[any])(nil)
