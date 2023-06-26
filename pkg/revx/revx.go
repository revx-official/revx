package revx

// The revx application name.
// This variable will be set on compilation.
var RevxApp = "unknown"

// The revx model name.
// This variable will be set on compilation.
var RevxModel = "unknown"

// The revx version.
// This variable will be set on compilation.
var RevxVersion = "unknown"

// The current revx commit.
// This variable will be set on compilation.
var RevxCommit = "unknown"

type RevxInfo struct {
	App     string `yaml:"app" json:"app"`
	Model   string `yaml:"model" json:"model"`
	Version string `yaml:"version" json:"version"`
	Commit  string `yaml:"commit" json:"commit"`
}

func Default() *RevxInfo {
	info := RevxInfo{}

	info.App = RevxApp
	info.Model = RevxModel
	info.Version = RevxVersion
	info.Commit = RevxCommit

	return &info
}
