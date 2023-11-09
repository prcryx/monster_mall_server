package types

import "net/http"

type IRoute interface {
	GetPath() string
	GetHandler() http.Handler
}

type Route struct {
	path    string
	handler http.Handler
}

func NewRoute(path string, handler http.Handler) Route {
	return Route{
		path:    path,
		handler: handler,
	}
}

var _ IRoute = (*Route)(nil)

func (r Route) GetPath() string {
	return r.path
}

func (r Route) GetHandler() http.Handler {
	return r.handler
}
