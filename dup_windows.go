//go:build windows

package cmdio

import (
	"syscall"

	"golang.org/x/sys/windows"
)

type Handle = syscall.Handle

func sys_dup(oldfd Handle) (Handle, error) {
	currentProcess := windows.CurrentProcess()
	var newFdHandle windows.Handle

	var hSourceProcessHandle = currentProcess
	var hSourceHandle = windows.Handle(oldfd)
	var hTargetProcessHandle = currentProcess
	var lpTargetHandle = &newFdHandle
	var dwDesiredAccess uint32 = 0
	var bInheritHandle = false
	var dwOptions uint32 = windows.DUPLICATE_SAME_ACCESS

	err := windows.DuplicateHandle(
		hSourceProcessHandle,
		hSourceHandle,
		hTargetProcessHandle,
		lpTargetHandle,
		dwDesiredAccess,
		bInheritHandle,
		dwOptions,
	)
	if err != nil {
		return 0, err
	}

	return Handle(newFdHandle), nil
}
