package golang

// Plugin is a type of protoc plugin you can use
type Plugin = string

// Path specifies how the paths of generated files are structured
type Path = string

// Plugins are all available plugins
var Plugins = []Plugin{GRPC}

// Paths are all available path options
var Paths = []Path{Import, SourceRelative}

const (
	// GRPC represents the plugin available to generate code for gRPC
	GRPC = Plugin("grpc")

	// Import represents a method for allowing your import paths to be as defined in the option go_package
	Import = Path("import")

	// SourceRelative represents a method for allowing your import paths to be defined by how they are organised in your protos directory
	SourceRelative = Path("source_relative")
)

// Config is configuration specific for the programming language go
type Config struct {
	Plugins []Plugin
	Paths   Path
}
