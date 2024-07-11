//go:build !windows

package cmdio

import (
	"syscall"
)

type Handle = int

func sys_dup(oldfd Handle) (Handle, error) {
	return syscall.Dup(oldfd)
}
