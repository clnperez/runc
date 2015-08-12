// +build linux,!go1.5 linux,gccgo

package libcontainer

import "syscall"

// GidMappingsEnableSetgroups was added in Go 1.5, so do nothing when building
// with earlier versions
func enableSetgroups(sys *syscall.SysProcAttr) {
}
