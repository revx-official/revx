package api

import (
	"github.com/revx-official/output/log"

	"github.com/revx-official/revx/pkg/config"
	"github.com/revx-official/revx/pkg/health"
	"github.com/revx-official/revx/pkg/proxy"
)

// Description:
//
//	Creates an endpoint for a route corresponding to the given http method.
//
// Parameters:
//
//	prox 	The reverse proxy.
//	method 	The http method to handle.
func CreateEndpointProxyHandler(prox *proxy.ReverseProxyServerInfo, method string) {
	handler := proxy.LoadBalancingHandler(prox)

	Router.ProxyHandle(method, prox.Context, handler)
	Router.ProxyHandle(method, prox.Context+"/*path", handler)
}

// Description:
//
//	Creates all specified endpoints for the given route.
//
// Parameters:
//
//	prox The reverse proxy.
func CreateEndpointsForProxy(prox *proxy.ReverseProxyServerInfo) {
	for _, method := range prox.AllowedMethods {
		CreateEndpointProxyHandler(prox, method)
	}

	healthCheck := health.NewHealthCheckRoutine(prox)
	health.RunHealthCheckRoutine(healthCheck)
}

// Description:
//
//	Creates the corresponding api endpoints for all registered proxy services.
func CreateReverseProxies() {
	for _, server := range config.Global.Servers {
		prox, err := proxy.NewReverseProxyServer(server)

		if err != nil {
			log.Fatalf("api: unable to create reverse proxy: %s", prox.Name)
		}

		CreateEndpointsForProxy(prox)
	}
}

// Description:
//
//	Initializes the reverse proxy api using the global configuration.
//	Provides the endpoints.
func InitProxyApi() {
	CreateReverseProxies()
}
