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
		elog.Info(fmt.Sprintf("Current Hostscan Version: %s", vars.VersionInfo))
		vars.ProcessBar.Clear()
		return
	}

	elog.Info("Hostscan Start! Waiting for your good news...")

	taskType := core.GetTaskType()
	if taskType == "noip"{
		elog.Error("No IP Found! Please use -i/-I to input single ip or ips in file")
		vars.ProcessBar.Clear()
		return
	}else if taskType == "nohost"{
		elog.Error("No Host Found! Please use -d/-D to input single host or hosts in file")
		vars.ProcessBar.Clear()
		return
	}

	if len(*vars.OutFile) > 0{
		exist,_ := utils.PathExists(*vars.OutFile)
		if exist{
			_ = os.Remove(*vars.OutFile)
		}
	}

	utils.SetUlimitMax()
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
