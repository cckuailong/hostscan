// +build windows

package utils

import "hostscan/elog"

func SetUlimitMax() {
	elog.Info("ignore ulimit on win os.")
}
