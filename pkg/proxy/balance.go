package proxy

import (
	"sync"
)

// Description:
//
//	The load balancer information is used by any proxy request handler.
//	It provides information used to select the proxy, which is going to handle the request.
type LoadBalancerInfo struct {
	ProxyInstanceIndex uint32     `json:"proxyInstanceIndex"` // The currently active proxy instance index.
	Mutex              sync.Mutex `json:"-"`                  // The mutex used to lock operations on the index.
}

// Description.
//
//	Creates a new load balancer info struct.
//
// Returns:
//
//	The create load balancer info.
func NewLoadBalancerInfo() *LoadBalancerInfo {
	return &LoadBalancerInfo{}
}
