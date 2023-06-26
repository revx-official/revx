package boot

import (
	"github.com/revx-official/output/log"
	"github.com/revx-official/revx/pkg/api"
	"github.com/revx-official/revx/pkg/config"
)

// Description:
//
//	Initializes the reverse proxy.
//	Reads the global configuration file and starts the api.
//
// Parameters:
//
//	configFilePath The configuration file path.
func Boot(configFilePath string) {
	log.Infof("boot: running revx ...")
	log.Infof("boot: reading configuration file ...")

	err := config.LoadConfig(configFilePath)

	if err != nil {
		log.Warnf("%s: %s", "boot: unable to load configuration file", err)
		log.Warnf("boot: falling back to default configuration ...")
	}

	api.InitApi()
	api.InitRevxApi()
	api.InitProxyApi()

	api.Boot()
}
