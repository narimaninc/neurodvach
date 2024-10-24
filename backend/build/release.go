//go:build !debug
// +build !debug

package build

const (
	CurrentConfig = DebugConfig

	Release = true
	Debug   = false
)
