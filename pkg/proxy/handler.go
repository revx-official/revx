package proxy

import (
	"net/http"

	"github.com/revx-official/output/log"
	"github.com/revx-official/revx/pkg/router"
)

// Description:
//
//	Represents the endpoint handler for any reverse proxy endpoint.
//	This handler performs round robin load balancing across the proxy instances.
//
// Parameters:
//
//	prox The reverse proxy.
func LoadBalancingHandler(prox *ReverseProxyServerInfo) router.RouterProxyHandlerFunc {
	return func(request *http.Request, response http.ResponseWriter) {
		log.Infof("proxy: pass %s %s", request.Method, request.URL.Path)

		prox.BalancerInfo.Mutex.Lock()

		instanceCount := uint32(len(prox.Upstreams))
		proxyIndex := prox.BalancerInfo.ProxyInstanceIndex % instanceCount

		instance := prox.Upstreams[proxyIndex]

		for !instance.HealthStats.Healthy {
			proxyIndex = (proxyIndex + 1) % instanceCount
			instance = prox.Upstreams[proxyIndex]
		}

		prox.BalancerInfo.ProxyInstanceIndex = (proxyIndex + 1) % instanceCount
		prox.BalancerInfo.Mutex.Unlock()

		instance.ReverseProxy.ServeHTTP(response, request)
	}
}
