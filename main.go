package main

import (
	"hostscan/core"
	"hostscan/elog"
	"hostscan/utils"
	"hostscan/vars"
	"flag"
	"fmt"
)


func main(){
	utils.Banner()
	flag.Parse()
	if *vars.Version {
		elog.Info(fmt.Sprintf("Current hostscan version: %s", vars.VersionInfo))
		return
	}

	if len(*vars.Ip) > 0 {
		vars.Ips = append(vars.Ips, *vars.Ip)
	}else if len(*vars.IpFile) > 0 {
		tmp_ips := utils.Readlines(*vars.IpFile)
		for _, tmp_ip := range tmp_ips{
			vars.Ips = append(vars.Ips, tmp_ip)
		}
	}

	if len(vars.Ips) == 0 {
		elog.Error("No IP Found! Please use -i/-I to input single ip or ips in file")
		return
	}

	if len(*vars.Host) > 0 {
		vars.Hosts = append(vars.Hosts, *vars.Host)
	}else if len(*vars.HostFile) > 0 {
		tmp_hosts := utils.Readlines(*vars.HostFile)
		for _, tmp_host := range tmp_hosts{
			vars.Hosts = append(vars.Hosts, tmp_host)
		}
	}

	if len(vars.Hosts) == 0 {
		elog.Error("No Host Found! Please use -d/-D to input single host or hosts in file")
		return
	}

	core.Scan()

	vars.ProcessBar.Finish()
}
