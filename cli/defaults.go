package cli

var (
	// GitCommit is the git commit hash of the source-tree during build
	GitCommit string

	// DefaultWorkDir is the default working directory
	// Set for OSX and Linux in code but can be overwritten with ldflags at build
	DefaultWorkDir = "/usr/local/var/protogen"
)
