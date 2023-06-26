package main

import (
	"flag"

	"github.com/revx-official/output/log"
	"github.com/revx-official/revx/pkg/boot"
	"github.com/revx-official/revx/pkg/config"
)

// Description:
//
//	Initializes this package.
//	Keep logging moderate for release builds.
func init() {
	log.Level = log.LevelInfo
}

// Specifies whether to enable verbose logging.
var flagVerbose bool

// Specifies the configuration file path.
var flagConfigFilePath string

// Description:
//
//	The application entry point.
func main() {
	flag.BoolVar(&flagVerbose, "verbose", false, "Enable trace logging.")
	flag.StringVar(&flagConfigFilePath, "config", config.DefaultConfigFilePath, "Specify the configuration file path.")

	flag.Parse()

	if flagVerbose {
		log.Level = log.LevelTrace
	}

	boot.Boot(flagConfigFilePath)
}
