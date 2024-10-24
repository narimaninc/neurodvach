package build

type BuildConfig uint8

const (
	ReleaseConfig BuildConfig = iota
	DebugConfig
)

func (c BuildConfig) String() string {
	switch c {
	case ReleaseConfig:
		return "release"
	case DebugConfig:
		return "debug"
	default:
		return "INVALID"
	}
}
