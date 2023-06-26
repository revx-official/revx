package api

import (
	"github.com/revx-official/output/log"
	"github.com/revx-official/revx/pkg/config"
	"github.com/revx-official/revx/pkg/router"
)

// The global router engine.
var Router router.Router

// Description:
//
//	Initializes the global router engine.
func InitApi() {
	Router = router.Default()
}

// Description:
//
//	Runs the global router engine, i.e. provides the api endpoints.
func Boot() {
	log.Infof("api: running server ...")
	log.Infof("api: serving on port: %d", config.Global.Port)

	err := Router.Run(config.Global.Port)

	if err != nil {
		log.Fatalf("api: unable to run server: %s", err)
	}

}
