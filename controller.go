package ring

import "context"

// A Controller handles the incoming request model, and returns a response
// model or error.
type Controller[Req, Res any] interface {
    Execute(context.Context, Req) (Res, error)
}
