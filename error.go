package ring

import "net/http"

// ErrorHandler provides a global error filter for all handlers
type ErrorHandler interface {
    Handle(w http.ResponseWriter, err error)
}
