package health

import (
	"sync"

	"github.com/revx-official/output/log"
)

type HeathCheckManager struct {
	Routines map[string]*HealthCheckRoutine
}

var Manager = &HeathCheckManager{
	Routines: make(map[string]*HealthCheckRoutine),
}

var mutex = sync.Mutex{}

func RegisterHealthCheckRoutine(healthCheck *HealthCheckRoutine) {
	mutex.Lock()
	defer mutex.Unlock()

	name := healthCheck.Proxy.Name
	_, exists := Manager.Routines[name]

	if exists {
		log.Warnf("health: routine already exists for proxy: %s", name)
		return
	}

	Manager.Routines[name] = healthCheck
}
