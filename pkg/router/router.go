package router

import "net/http"

// Description:
//
//	Function definition for router endpoint handlers.
type RouterHandlerFunc = func(request *Request) *Response

// Description:
//
//	Function definition for router endpoint handlers.
type RouterProxyHandlerFunc = func(request *http.Request, response http.ResponseWriter)

// Description:
//
//	A router request.
//	Represents a HTTP request.
type Request struct {
	Method          string
	Url             string
	Path            string
	Headers         map[string]string `json:"headers"`
	PathParameters  map[string]string `json:"pathParameters"`
	QueryParameters map[string]string `json:"queryParameters"`
	Body            string            `json:"body"`
}

// Description:
//
//	A router response.
//	Represents a HTTP response.
type Response struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       interface{}       `json:"body"`
}

// Description:
//
//	The router interface.
type Router interface {
	Handle(method string, path string, handler RouterHandlerFunc)
	ProxyHandle(method string, path string, handler RouterProxyHandlerFunc)
	Run(port uint16) error
}

// Description:
//
//	Creates the default router.
//
// Returns:
//
//	The default router.
func Default() Router {
	return NewGinRouter()
}
