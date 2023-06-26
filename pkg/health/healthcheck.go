package health

import (
	"net/http"
	"time"

	"github.com/revx-official/output/log"
	"github.com/revx-official/revx/pkg/proxy"
)

type HealthCheckRoutine struct {
	Proxy  *proxy.ReverseProxyServerInfo
	Ticker *time.Ticker
	Cancel chan bool
}

func NewHealthCheckRoutine(proxy *proxy.ReverseProxyServerInfo) *HealthCheckRoutine {
	healthCheck := HealthCheckRoutine{}

	interval := time.Duration(proxy.HealthCheckInfo.Interval)

	healthCheck.Proxy = proxy
	healthCheck.Ticker = time.NewTicker(interval * time.Millisecond)
	healthCheck.Cancel = make(chan bool)

	RegisterHealthCheckRoutine(&healthCheck)
	return &healthCheck
}

func RunHealthCheckRoutine(healthCheck *HealthCheckRoutine) {
	go internalRunHealthCheckRoutine(healthCheck)
}

func internalRunHealthCheckRoutine(healthCheck *HealthCheckRoutine) {
	for {
		select {
		case timeStamp := <-healthCheck.Ticker.C:
			healthCheckRoutineInterval(healthCheck, timeStamp)
		case <-healthCheck.Cancel:
			return
		}
	}
}

func healthCheckRoutineInterval(healthCheck *HealthCheckRoutine, timeStamp time.Time) {
	for index := range healthCheck.Proxy.Upstreams {
		instanceRef := &healthCheck.Proxy.Upstreams[index]

		log.Tracef("health: running check: get %s %s", instanceRef.TargetUrl.String(), timeStamp)
		response, err := http.Get(instanceRef.TargetUrl.String())

		if err != nil {
			instanceRef.HealthStats.Error = err.Error()
			instanceRef.HealthStats.ConsecutiveFails = instanceRef.HealthStats.ConsecutiveFails + 1

			log.Tracef("health: check failed: %s, consecutive fails: %d", err, instanceRef.HealthStats.ConsecutiveFails)

			if instanceRef.HealthStats.ConsecutiveFails >= healthCheck.Proxy.HealthCheckInfo.Fails {
				instanceRef.HealthStats.Healthy = false
				log.Tracef("health: service unhealthy: %s", instanceRef.TargetUrl.String())
			}

			continue
		}

		defer response.Body.Close()

		instanceRef.HealthStats.Healthy = true
		instanceRef.HealthStats.ConsecutiveFails = 0
		instanceRef.HealthStats.Error = ""

		log.Tracef("health: check succeeded: get %s, status: %d", instanceRef.TargetUrl.String(), response.StatusCode)
	}
}
