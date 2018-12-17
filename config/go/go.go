package golang

import "fmt"

// Plugin is a type of protoc plugin you can use
type Plugin = string

// Path specifies how the paths of generated files are structured
type Path = string

const (
	// GRPC represents the plugin available to generate code for gRPC
	GRPC = Plugin("grpc")

	// Import represents a method for allowing your import paths to be as defined in the option go_package
	Import = Path("import")

	// SourceRelative represents a method for allowing your import paths to be defined by how they are organised in your protos directory
	SourceRelative = Path("source_relative")
)

// Plugins are all available plugins
var Plugins = []Plugin{GRPC}

// AllowedPlugins is a whitelist for checking for allowed plugins
var AllowedPlugins = map[Plugin]struct{}{
	GRPC: struct{}{},
}

// ErrPluginNotAllowed happens when the plugin does not exist
type ErrPluginNotAllowed struct {
	plugin Plugin
}

func (e ErrPluginNotAllowed) Error() string {
	return fmt.Sprintf("go plugin does not exist: %q", e.plugin)
}

// IsAllowedPlugin returns true if the plugin is allowed
func IsAllowedPlugin(plugin Plugin) error {
	if _, ok := AllowedPlugins[plugin]; ok {
		return ErrPluginNotAllowed{plugin}
	}
	return nil
}

// Paths are all available path options
var Paths = []Path{Import, SourceRelative}

// AllowedPaths is a whitelist for checking for allowed type of pathing
var AllowedPaths = map[Path]struct{}{
	Import:         struct{}{},
	SourceRelative: struct{}{},
}

// ErrPathNotAllowed happens when the pathing option is invalid
type ErrPathNotAllowed struct {
	path Path
}

func (e ErrPathNotAllowed) Error() string {
	return fmt.Sprintf("go pathing option is invalid: %q", e.path)
}

// IsAllowedPath returns true if the type of pathing is allowed
func IsAllowedPath(path Path) error {
	if _, ok := AllowedPaths[path]; ok {
		return ErrPathNotAllowed{path}
	}
	return nil
}

// Config is configuration specific for the programming language go
type Config struct {
	Plugins []Plugin
	Paths   Path
}
