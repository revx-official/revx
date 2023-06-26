package proxy

import (
	"sync"

	"github.com/revx-official/output/log"
)

// Description:
//
//	The proxy manager.
//	Stores all registered reverse proxies.
type ProxyManagerInfo struct {
	Proxies map[string]*ReverseProxyServerInfo `json:"proxies"` // Contains all registered proxies.
}

// The global proxy manager instance.
var Manager = &ProxyManagerInfo{
	Proxies: make(map[string]*ReverseProxyServerInfo),
}

// The internal mutex to control access to the global manager.
var mutex = sync.Mutex{}

// Description:
//
//	Registers a proxy in the global proxy manager instance.
//
// Parameters:
//
//	proxy The reverse proxy to register.
func RegisterProxy(proxy *ReverseProxyServerInfo) {
	mutex.Lock()
	defer mutex.Unlock()

	_, exists := Manager.Proxies[proxy.Name]

	if exists {
		log.Warnf("proxy: proxy already exists for service: %s", proxy.Name)
		return
	}

	Manager.Proxies[proxy.Name] = proxy
}
