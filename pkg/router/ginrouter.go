package router

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// Description:
//
//	Implementation of the Router interface for gin.
type GinRouter struct {
	engine *gin.Engine
}

// Description:
//
//	Package initializer.
//	Sets gin to release mode.
func init() {
	gin.SetMode(gin.ReleaseMode)
}

// Description:
//
//	Creates a new gin router.
//
// Returns:
//
//	The created gin router.
func NewGinRouter() *GinRouter {
	engine := gin.New()

	engine.RedirectTrailingSlash = true
	engine.RedirectFixedPath = true

	return &GinRouter{
		engine: engine,
	}
}

// Description:
//
//	Registers a new HTTP handler function for the given method and path.
//	Paths can include wildcards and path variables.
//
// Parameters:
//
//	method 	The http method to handle.
//	path   	The path to handle.
//	handler	The handler responsible for handling the request.
func (router *GinRouter) Handle(method string, path string, handler RouterHandlerFunc) {
	router.engine.Handle(method, path, func(context *gin.Context) {
		internalRouteHandler(path, context, handler)
	})
}

// Description:
//
//	Registers a new HTTP handler function for the given method and path.
//	A registered handler is provided with the native request and response writer.
//	Paths can include wildcards and path variables.
//
// Parameters:
//
//	method 	The http method to handle.
//	path   	The path to handle.
//	handler	The handler responsible for handling the request.
func (router *GinRouter) ProxyHandle(method string, path string, handler RouterProxyHandlerFunc) {
	router.engine.Handle(method, path, func(context *gin.Context) {
		internalProxyRouteHandler(context, handler)
	})
}

// Description:
//
//	Starts the HTTP server for this router and listens to all registered routes.
//
// Returns:
//
//	An error if serving the router fails.
func (router *GinRouter) Run(port uint16) error {
	portFmt := fmt.Sprintf(":%d", port)

	// server := &http.Server{
	// 	Addr:    portFmt,
	// 	Handler: router.engine,
	// }

	// // Setting this to false apparently reduces memory usage.
	// // However, setting this to true apparently is the standard and improves performance.
	// server.SetKeepAlivesEnabled(true)

	// return server.ListenAndServe()

	return router.engine.Run(portFmt)
}

// Description:
//
//	Internal handler method for incoming requests.
//	Triggered by the gin framework.
//
// Parameters:
//
//	pathHandle 	The registered path handle.
//	context 	The internal gin context.
//	handler 	The registered handler function.
func internalRouteHandler(pathHandle string, context *gin.Context, handler RouterHandlerFunc) {
	request := context.Request

	internalRequest, err := transformRequest(pathHandle, request)

	if err != nil {
		panic("router: cannot transform request")
	}

	internalResponse := handler(internalRequest)
	applyResponse(internalResponse, context)
}

// Description:
//
//	Internal handler method for incoming requests.
//	Triggered by the gin framework.
//
// Parameters:
//
//	context 	The internal gin context.
//	handler 	The registered handler function.
func internalProxyRouteHandler(context *gin.Context, handler RouterProxyHandlerFunc) {
	handler(context.Request, context.Writer)
}

// Description:
//
//	Transforms an incoming HTTP request to a router request.
//
// Parameters:
//
//	pathHandle 	The registered path handle.
//	request		The request to transform.
//
// Returns:
//
//	The transformed request, or an error, if the request could not be transformed.
func transformRequest(pathHandle string, request *http.Request) (*Request, error) {
	result := Request{
		Method:          request.Method,
		Url:             request.URL.String(),
		Path:            request.URL.Path,
		Headers:         make(map[string]string),
		PathParameters:  make(map[string]string),
		QueryParameters: make(map[string]string),
	}

	for key, values := range request.Header {
		result.Headers[key] = strings.Join(values, ",")
	}

	pathParameters, err := extractPathParameters(pathHandle, request.URL.Path)
	if err != nil {
		return nil, err
	}

	result.PathParameters = pathParameters

	queryParameters, err := extractQueryParameters(request.URL.String())
	if err != nil {
		return nil, err
	}

	result.QueryParameters = queryParameters
	return &result, nil
}

// Description:
//
//	Extracts all path parameters using the registered path handle and the actual request path.
//
// Example:
//   - handle: 	/some/path/:variable
//   - path:	/some/path/128
//
// Parameters:
//
//	handle The registered path handle.
//	path The actual request path.
func extractPathParameters(handle string, path string) (map[string]string, error) {
	result := make(map[string]string)

	path = strings.TrimPrefix(path, "/")
	handle = strings.TrimPrefix(handle, "/")

	pathSegments := strings.Split(path, "/")
	handleSegments := strings.Split(handle, "/")

	if len(handleSegments) != len(pathSegments) {
		return nil, fmt.Errorf("router: number of url segments does not match number of path segments")
	}

	for index, segment := range handleSegments {
		if !strings.HasPrefix(segment, ":") {
			continue
		}

		paramName := strings.TrimPrefix(segment, ":")
		result[paramName] = pathSegments[index]
	}

	return result, nil
}

// Description:
//
//	Extracts all query parameters from the request path.
//
// Parameters:
//
//	path The actual request path.
func extractQueryParameters(path string) (map[string]string, error) {
	parameters := make(map[string]string)

	parsedURL, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	query := parsedURL.Query()

	for key, values := range query {
		parameters[key] = strings.Join(values, ",")
	}

	return parameters, nil
}

// Description:
//
//	Applies a router response to the internal gin context.
func applyResponse(response *Response, context *gin.Context) {
	for key, value := range response.Headers {
		context.Header(key, value)
	}

	context.JSON(response.StatusCode, response.Body)
}
