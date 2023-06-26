package api

import (
	"net/http"

	"github.com/revx-official/output/log"
	"github.com/revx-official/revx/pkg/config"
	"github.com/revx-official/revx/pkg/proxy"
	"github.com/revx-official/revx/pkg/revx"
	"github.com/revx-official/revx/pkg/router"
)

// Description:
//
//	Represents an error response for the info api.
type InfoErrorResponse struct {
	Message string `json:"message"` // The error message.
}

// Description:
//
//	Endpoint: /info
//
// Parameters:
//
//	context The http context.
func HandleInfo(request *router.Request) *router.Response {
	log.Infof("%s: %s", "api: request", request.Path)

	return &router.Response{
		StatusCode: http.StatusOK,
		Body:       revx.Default(),
	}
}

// Description:
//
//	Endpoint: /config
//
// Parameters:
//
//	context The http context.
func HandleConfig(request *router.Request) *router.Response {
	log.Infof("%s: %s", "api: request", request.Path)

	return &router.Response{
		StatusCode: http.StatusOK,
		Body:       config.Global,
	}
}

// Description:
//
//	Endpoint: /inspect
//
// Parameters:
//
//	context The http context.
func HandleInspect(request *router.Request) *router.Response {
	log.Infof("%s: %s", "api: request", request.Path)

	return &router.Response{
		StatusCode: http.StatusOK,
		Body:       proxy.Manager,
	}
}

// Description:
//
//	Endpoint: /inspect/:name
//
// Parameters:
//
//	context The http context.
func HandleInspectByName(request *router.Request) *router.Response {
	log.Infof("%s: %s", "api: request", request.Path)

	name := request.PathParameters["name"]
	prox := proxy.Manager.Proxies[name]

	if prox == nil {
		return &router.Response{
			StatusCode: http.StatusBadRequest,
			Body:       InfoErrorResponse{Message: "Proxy not found."},
		}
	}

	return &router.Response{
		StatusCode: http.StatusOK,
		Body:       prox,
	}
}

// Description:
//
//	Initializes the development/info api.
//	This api is used to retrieve internal information about the revx proxy service.
func InitRevxApi() {
	Router.Handle("GET", "revx/info", HandleInfo)
	Router.Handle("GET", "revx/config", HandleConfig)

	Router.Handle("GET", "revx/inspect", HandleInspect)
	Router.Handle("GET", "revx/inspect/:name", HandleInspectByName)
}
