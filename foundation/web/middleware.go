package web

// Middleware is a funcion that is used to run some code between the request response cycle
// Here, Middleware is a type that takes in a Handler and returns a Handler wrapped around
// with some code
type Middleware func(Handler) Handler

// WrapMiddleware wraps middlewares around a handler and returns a handler
// The middlewares are execuuted in the order they are in the slice
func WrapMiddleware(middlewares []Middleware, handler Handler) Handler {

	for i := len(middlewares) - 1; i >= 0; i-- {
		h := middlewares[i]
		if h != nil {
			// Wrap the current handler with the middleware
			// which gives a new handler
			handler = h(handler)
		}
	}

	return handler
}
