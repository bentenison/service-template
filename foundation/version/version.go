package version

type VersionInfo struct {
	Version      string // Application version
	GitCommit    string // Git commit hash
	BuildTime    string // Time of the build
	GoVersion    string // Go version used
	TargetOS     string // Operating System
	TargetArch   string // System Architecture
	ContainerId  string
	LanguageName string
}

var BuildInfo VersionInfo
