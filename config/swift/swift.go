package swift

// FileNaming represents the option for file naming of generated sources
type FileNaming = string

const (
	// FileNamingName is the option name for file naming
	FileNamingName = "FileNaming"

	// FullPath is the default and maps the exact path of the .proto file
	FullPath = FileNaming("FullPath")

	// PathToUnderscores compress the path to use underscores instead of slashes
	PathToUnderscores = FileNaming("PathToUnderscores")

	// DropPath doesn't include a path and only saves file to the output
	// directory with the same filename as the .proto file
	DropPath = FileNaming("DropPath")
)

// FileNamingValues represents the actual values for file naming
var FileNamingValues = map[FileNaming]struct{}{
	FullPath:          struct{}{},
	PathToUnderscores: struct{}{},
	DropPath:          struct{}{},
}

// Visibility represents the option for visibility of generated types
type Visibility = string

const (
	// VisibilityName is the option name for generated class visibility
	VisibilityName = "Visibility"

	// Internal visibility of generated types
	Internal = Visibility("Internal")

	// Public visibility of generated types
	Public = Visibility("Public")
)

// VisibilityValues represents the visibility values that are allowed
var VisibilityValues = map[Visibility]struct{}{
	Public:   struct{}{},
	Internal: struct{}{},
}

const (
	// ProtoPathModuleMappingsName is the option name for the swift module file
	ProtoPathModuleMappingsName = "ProtoPathModuleMappings"
)

// OptionNames represents the known option names
var OptionNames = map[string]struct{}{
	FileNamingName:              struct{}{},
	VisibilityName:              struct{}{},
	ProtoPathModuleMappingsName: struct{}{},
}

// Option represents a configuration option for swift
type Option struct {
	Name  string
	Value string
}

// Config is configuration specific for the programming language Swift
type Config struct {
	Options []Option
}
