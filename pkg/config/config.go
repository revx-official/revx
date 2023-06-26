package config

import (
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v3"
)

// Description:
//
//	Represents the top level configuration.
type ConfigRevx struct {

	// The port to run revx on.
	Port uint16 `yaml:"port" json:"port"`

	// The server configuration.
	Servers []ConfigReverseProxyServer `yaml:"servers" json:"servers,omitempty"`
}

// Description:
//
//	Represents a service configuration.
type ConfigReverseProxyServer struct {

	// The name of the server.
	// Only serves identification purposes.
	Name string `yaml:"name" json:"name"`

	// The context path which is used to route to the server.
	// All requested routes starting with this context path will be proxy forwarded to one of the given upstreams.
	// Context paths must be unique.
	//
	// If the context path of this server is example/,
	// the following routes will be forwarded to this server:
	//	- example/
	//	- example/value
	//	- example/subroute/value
	//	- example/*
	Context string `yaml:"context" json:"context"`

	// The registered server upstreams.
	// Any server can have multiple upstreams.
	// revx is then going ahead and load balances traffic between all registered upstreams.
	Upstreams []string `yaml:"upstreams" json:"upstreams"`

	// The allowed http methods, e.g. GET, POST, ...
	AllowedMethods []string `yaml:"allowed-methods" json:"allowedMethods"`

	// The health check configuration.
	HealthCheck ConfigReverseProxyServerHealthCheck `yaml:"health-check" json:"healthCheck"`
}

// Description:
//
// Represents a service health check configuration.
type ConfigReverseProxyServerHealthCheck struct {

	// The health check endpoint.
	Endpoint string `yaml:"endpoint" json:"endpoint"`

	// The health check interval.
	Interval uint32 `yaml:"interval" json:"interval"`

	// The maximum amount of fails.
	// If this amount of fails is exceeded, the upstream is considered unhealthy.
	Fails uint32 `yaml:"fails" json:"fails"`
}

// The global configuration.
var Global = Default()

// The internal mutex used to control access to the global configuration.
var mutex = sync.Mutex{}

// Description:
//
//	Creates the default reverse proxy configuration.
//
// Returns:
//
//	The default reverse proxy configuration.
func Default() *ConfigRevx {
	return &ConfigRevx{
		Port: DefaultPort,
	}
}

// Description:
//
//	Loads the configuration from a given file path.
//
// Parameters:
//
//	path The configuration file path.
//
// Returns:
//
//	The reverse proxy configuration object.
func LoadConfig(path string) error {
	mutex.Lock()
	defer mutex.Unlock()

	file, err := readConfigFile(path)

	if err != nil {
		return err
	}

	config, err := unmarshalConfig(file)

	if err != nil {
		return err
	}

	Global = config
	return nil
}

// Description:
//
//	Reads the global configuration file.
//
// Parameters:
//
//	path The path of the configuration file.
//
// Returns:
//
//	The content of the global configuration file in bytes.
func readConfigFile(path string) ([]byte, error) {
	config, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return config, nil
}

// Description:
//
//	Unmarshals the configuration file content, given in bytes.
//
// Parameters:
//
//	config The configuration file content in bytes.
//
// Returns:
//
//	The reverse proxy configuration object.
func unmarshalConfig(file []byte) (*ConfigRevx, error) {
	config := ConfigRevx{}
	err := yaml.Unmarshal(file, &config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}
