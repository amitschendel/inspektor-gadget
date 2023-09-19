package syscalls

import (
	"runtime"

	amd64syscalls "github.com/inspektor-gadget/inspektor-gadget/pkg/utils/syscalls/amd64"
	arm64syscalls "github.com/inspektor-gadget/inspektor-gadget/pkg/utils/syscalls/arm64"
)

func GetSyscallNumberByName(arch string, name string) (int, bool) {
	if arch == "" {
		arch = runtime.GOARCH
	}

	var number int
	var ok bool

	if arch == "amd64" {
		number, ok = amd64syscalls.SyscallsNameToNumber[name]
	} else if arch == "arm64" {
		number, ok = arm64syscalls.SyscallsNameToNumber[name]
	} else {
		return -1, false
	}

	return number, ok
}

func GetSyscallNameByNumber(arch string, number int) (string, bool) {
	if arch == "" {
		arch = runtime.GOARCH
	}

	var name string
	var ok bool

	if arch == "amd64" {
		name, ok = amd64syscalls.SyscallsNumberToName[number]
	} else if arch == "arm64" {
		name, ok = arm64syscalls.SyscallsNumberToName[number]
	} else {
		return "", false
	}

	return name, ok
}
