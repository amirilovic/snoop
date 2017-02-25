package http

import "github.com/julienschmidt/httprouter"

// Route defines a route
type Route struct {
	Method string
	Path   string
	Handle httprouter.Handle
}
