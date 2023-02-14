package main

import (
	"flag"
	"fmt"
	"hostscan/core"
	"hostscan/elog"
	"hostscan/utils"
	"hostscan/vars"
	"os"
)

func main(){
	utils.Banner()

	flag.Parse()
	if *vars.Version {
		elog.Info(fmt.Sprintf("Current hostscan version: %s", vars.VersionInfo))
		return
	}

	utils.SetUlimitMax()

	taskType := core.GetTaskType()
	if taskType == "noip"{
		elog.Error("No IP Found! Please use -i/-I to input single ip or ips in file")
		return
	}else if taskType == "nohost"{
		elog.Error("No Host Found! Please use -d/-D to input single host or hosts in file")
		return
	}

	if len(*vars.OutFile) > 0{
		exist,_ := utils.PathExists(*vars.OutFile)
		if exist{
			_ = os.Remove(*vars.OutFile)
		}
	}

	err := core.Scan(taskType)
	if err != nil {
		elog.Error(fmt.Sprintf("Scan Failed: %v", err))
	}

	err = vars.ProcessBar.Finish()
	if err != nil {
		elog.Error(fmt.Sprintf("ProcessBar Close Failed: %v", err))
		return
	}
}
