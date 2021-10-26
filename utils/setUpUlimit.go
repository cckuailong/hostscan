// +build linux darwin

package utils

import (
	"fmt"
	"hostscan/elog"
	"syscall"
)

func SetUlimitMax() {
	var rLimit syscall.Rlimit

	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		elog.Warn(fmt.Sprintf("Warn Getting Rlimit %v", err.Error()))
		return
	}
	rLimit.Max = 10240
	rLimit.Cur = 10240
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		elog.Warn(fmt.Sprintf("Warn Setting Rlimit %v", err.Error()))
	}
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		elog.Warn(fmt.Sprintf("Warn Getting Rlimit %v", err.Error()))
	}
	elog.Info(fmt.Sprintf("Rlimit set success! %v", rLimit))
}

