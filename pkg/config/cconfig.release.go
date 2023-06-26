//go:build !local

package config

// Constant declarations.
const (
	// The default port.
	DefaultPort uint16 = 80

	// The default configuration file path.
	DefaultConfigFilePath string = "/etc/revx/config.yaml"
)
