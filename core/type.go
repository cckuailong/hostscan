package core

import (
	"fmt"
	"hostscan/vars"
	"net"
	"sort"
	"strconv"
	"strings"
)

const PORTMAX = 65535
const PORTMIN = 1

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

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func HandleIpRange(ipRange string)  []string{
	ip, ipNet, err := net.ParseCIDR(ipRange)
	if err != nil {
		return []string{}
	}

	var ips []string
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	// Remove network address and broadcast address
	return ips[1 : len(ips)-1]
}

func parsePort(ports string) []int {
	var scanPorts []int
	slices := strings.Split(ports, ",")
	var start_str, end_str string
	for _, port := range slices {
		port = strings.Trim(port, " ")
		if len(port) == 0{
			continue
		}
		if strings.Contains(port, "-") {
			ranges := strings.Split(port, "-")
			if len(ranges) < 2 {
				continue
			}
			sort.Strings(ranges)
			start_str = ranges[0]
			end_str = ranges[1]
			start, err := strconv.Atoi(start_str)
			if err != nil{
				continue
			}
			end, err := strconv.Atoi(end_str)
			if err != nil{
				continue
			}
			if start < PORTMIN{
				start = PORTMIN
			}
			if end > PORTMAX{
				end = PORTMAX
			}
			for i := start; i <= end; i++ {
				scanPorts = append(scanPorts, i)
			}
		}else{
			target_port, err := strconv.Atoi(port)
			if err == nil{
				scanPorts = append(scanPorts, target_port)
			}
		}

	}
	return scanPorts
}

type TaskInput struct {
	Host string
	IP   string
}

func HandleCustomPorts(host, ip string) []TaskInput{
	handled_set := []TaskInput{}
	clear_ip := strings.Split(ip, ":")[0]

	iports := parsePort(*vars.Iports)

	if len(iports) > 0 {
		for _,iport := range iports{
			handled_set = append(handled_set, TaskInput{
				Host: host,
				IP: fmt.Sprintf("%s:%d", clear_ip, iport),
			})
		}
	}else{
		handled_set = append(handled_set, TaskInput{
			Host: host,
			IP: ip,
		})
	}

	//fmt.Println(handled_set)

	return handled_set
}

