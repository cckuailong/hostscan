package core

import (
	"hostscan/vars"
)



func GetTaskType() string{
	if len(*vars.Ip) == 0 && len(*vars.IpFile) == 0{
		return "noip"
	}
	if len(*vars.Host) == 0 && len(*vars.HostFile) == 0{
		return "nohost"
	}


	if len(*vars.Ip) > 0{
		if len(*vars.Host) > 0{
			return "ip_host"
		}else if len(*vars.HostFile) > 0{
			return "ip_hostfile"
		}
	}else if len(*vars.IpFile) > 0{
		if len(*vars.Host) > 0{
			return "ipfile_host"
		}else if len(*vars.HostFile) > 0{
			return "ipfile_hostfile"
		}
	}

	return "N/A"
}
