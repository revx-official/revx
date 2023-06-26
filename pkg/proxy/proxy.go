package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/revx-official/revx/pkg/config"
)

// Description:
//
//	Represents a reverse proxy. A reverse proxy can consist of multiple reverse proxy instances,
//	where every instance represents exactly one server.
//	A reverse proxy can load balance across multiple instances of the same service.
type ReverseProxyServerInfo struct {
	Name            string                            `json:"name"`            // The name of the proxy.
	Context         string                            `json:"context"`         // The context path.
	AllowedMethods  []string                          `json:"allowedMethods"`  // All allowed http methods.
	Upstreams       []ReverseProxyServerUpstreamInfo  `json:"upstreams"`       // The individual reverse proxy instances.
	HealthCheckInfo ReverseProxyServerHealthCheckInfo `json:"healthCheckInfo"` // The reverse proxy health check information.
	BalancerInfo    LoadBalancerInfo                  `json:"balancerInfo"`    // Information used by the load balancer.
}

// Description:
//
//	Represents a single reverse proxy instance.
//	A reverse proxy instance proxy passes all requests to exactly one target server,
//	represented by the target url.
//	Each instance stores a bunch of stats related to health checks.
type ReverseProxyServerUpstreamInfo struct {
	TargetUrl    *url.URL                              `json:"targetUrl"`   // The url which is targeted by the reverse proxy.
	ReverseProxy *httputil.ReverseProxy                `json:"-"`           // The http reverse proxy.
	HealthStats  ReverseProxyServerUpstreamHealthStats `json:"healthStats"` // The instance health stats.
	Stats        ReverseProxyServerUpstreamStats       `json:"stats"`       // The instance statistics.
}

// Description:
//
//	Holds information about the health checks, which are executed on a proxy.
//	Health checks can be performed on a specific endpoint, if the target service
//	wants to track some information about its health check requests.
type ReverseProxyServerHealthCheckInfo struct {
	Endpoint string `json:"endpoint"` // The health check endpoint.
	Interval uint32 `json:"interval"` // The health check interval.
	Fails    uint32 `json:"fails"`    // The maximum amount of fails until a service is considered as unhealthy.
}

// Description:
//
//	Holds information about the health state of a single reverse proxy instance.
type ReverseProxyServerUpstreamHealthStats struct {
	Healthy          bool   `json:"healthy"`          // Whether the service is healthy.
	ConsecutiveFails uint32 `json:"consecutiveFails"` // The amount of consecutive health check fails experienced with this service.
	Error            string `json:"error"`            // The error, if there is one.
}

// Descriptions:
//
//	Tracks some statistics about a server upstream.
type ReverseProxyServerUpstreamStats struct {
	AverageRequestTime float32 `json:"averageRequestTime"`
}

// Description:
//
//	Creates a new reverse proxy based on the given configuration.
//	Registers the proxy in the global proxy manager.
//
// Parameters:
//
//	conf The reverse proxy configuration
//
// Returns:
//
//	The created reverse proxy.
func NewReverseProxyServer(conf config.ConfigReverseProxyServer) (*ReverseProxyServerInfo, error) {
	proxy := ReverseProxyServerInfo{}

	proxy.Name = conf.Name
	proxy.Context = conf.Context
	proxy.AllowedMethods = conf.AllowedMethods

	proxy.HealthCheckInfo.Endpoint = conf.HealthCheck.Endpoint
	proxy.HealthCheckInfo.Interval = conf.HealthCheck.Interval
	proxy.HealthCheckInfo.Fails = conf.HealthCheck.Fails

	for _, upstream := range conf.Upstreams {
		instance, err := NewReverseProxyServerUpstream(upstream)

		if err != nil {
			return nil, err
		}

		proxy.Upstreams = append(proxy.Upstreams, *instance)
	}

	RegisterProxy(&proxy)
	return &proxy, nil
}

// Description:
//
//	Creates a new reverse proxy.
//	The reverse proxy is initialized with default health stats.
//
// Parameters:
//
//	remote The url which should be targeted by the reverse proxy.
func NewReverseProxyServerUpstream(remote string) (*ReverseProxyServerUpstreamInfo, error) {
	target, err := url.Parse(remote)

	if err != nil {
		return nil, err
	}

	upstream := ReverseProxyServerUpstreamInfo{}

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.Transport = NewReverseProxyTransport(&upstream, http.DefaultTransport)

	healthStats := ReverseProxyServerUpstreamHealthStats{
		Healthy:          true,
		ConsecutiveFails: 0,
	}

	stats := ReverseProxyServerUpstreamStats{
		AverageRequestTime: 0,
	}

	upstream.TargetUrl = target
	upstream.ReverseProxy = proxy
	upstream.HealthStats = healthStats
	upstream.Stats = stats

	return &upstream, nil
}
