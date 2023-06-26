package proxy

import (
	"net/http"
	"time"

	"github.com/revx-official/output/log"
)

type ReverseProxyTransport struct {
	ProxyInstance *ReverseProxyServerUpstreamInfo
	Transport     http.RoundTripper
}

func NewReverseProxyTransport(proxy *ReverseProxyServerUpstreamInfo, transport http.RoundTripper) ReverseProxyTransport {
	return ReverseProxyTransport{ProxyInstance: proxy, Transport: transport}
}

func (transport ReverseProxyTransport) RoundTrip(request *http.Request) (response *http.Response, err error) {
	start := time.Now()

	response, err = transport.Transport.RoundTrip(request)
	duration := time.Since(start)

	log.Tracef("proxy: pass info: %s %s", request.Method, request.URL, duration)

	milliseconds := float32(duration * time.Millisecond)
	transport.ProxyInstance.Stats.AverageRequestTime = (transport.ProxyInstance.Stats.AverageRequestTime + milliseconds) / 2.0

	return response, err
}
