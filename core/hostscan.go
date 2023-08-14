package core

import (
	"bufio"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"hostscan/elog"
	"hostscan/utils"
	"hostscan/vars"
	"io"
	"os"
	"strings"
	"sync"
)


func calcTaskTotal(taskType string, SingleIpCnt int) int{
	var err error
	var ipCnt, hostCnt, schemeCnt int
	schemeCnt = len(vars.Schemes)
	if taskType == "ip_host" {
		ipCnt = SingleIpCnt
		hostCnt = 1
	}else if taskType == "ipfile_host" {
		ipCnt, err = utils.LineCounter(*vars.IpFile)
		if err != nil{
			elog.Error(fmt.Sprintf("Get Lines Count[%s]: %v", *vars.IpFile, err))
			return 0
		}
		hostCnt = 1
	}else if taskType == "ip_hostfile" {
		ipCnt = SingleIpCnt
		hostCnt, err = utils.LineCounter(*vars.HostFile)
		if err != nil{
			elog.Error(fmt.Sprintf("Get Lines Count[%s]: %v", *vars.HostFile, err))
			return 0
		}
	}else if taskType == "ipfile_hostfile" {
		ipCnt, err = utils.LineCounter(*vars.IpFile)
		if err != nil{
			elog.Error(fmt.Sprintf("Get Lines Count[%s]: %v", *vars.IpFile, err))
			return 0
		}
		hostCnt, err = utils.LineCounter(*vars.HostFile)
		if err != nil{
			elog.Error(fmt.Sprintf("Get Lines Count[%s]: %v", *vars.HostFile, err))
			return 0
		}
	}else{
		return 0
	}

	totalTask := ipCnt * hostCnt * schemeCnt

	elog.Info(fmt.Sprintf("Total Task: %d   ||   Ip: %d, Host: %d, Scheme:%d", totalTask, ipCnt, hostCnt, schemeCnt))

	return totalTask
}

func Scan(taskType string) error{
	wg := sync.WaitGroup{}
	ip_list := []string{}
	if strings.Contains(*vars.Ip, "/"){
		ip_list = HandleIpRange(*vars.Ip)
	}else{
		ip_list = append(ip_list, *vars.Ip)
	}

	totalTask := calcTaskTotal(taskType, len(ip_list))

	if totalTask == 0{
		elog.Error(fmt.Sprintf("Get Lines Count: 0"))
		return nil
	}

	vars.ProcessBar = progressbar.NewOptions(totalTask,
		progressbar.OptionClearOnFinish(),
		progressbar.OptionEnableColorCodes(false),
		progressbar.OptionShowCount(),
		progressbar.OptionFullWidth(),
		progressbar.OptionSetDescription("[*] Scanning..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "=",
			SaucerHead:    ">",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	// 创建一个buffer为vars.ScanNum * 4的channel
	taskChan := make(chan Task, *vars.Thread*4)

	// 创建vars.ThreadNum个协程
	for i := 0; i < *vars.Thread; i++ {
		go goScan(taskChan, &wg)
		wg.Add(1)
	}

	if taskType == "ip_host" {
		for _,ip := range ip_list{
			for _, scheme := range vars.Schemes {
				handled_set := HandleCustomPorts(*vars.Host, ip)
				for _,item := range handled_set{
					task := Task{
						Uri:  fmt.Sprintf("%s://%s", scheme, item.IP),
						Host: item.Host,
					}
					// 生产者，不断地往taskChan channel发送数据，直到channel阻塞
					taskChan <- task
				}
			}
		}

	}else if taskType == "ipfile_host" {
		ip_f, err := os.Open(*vars.IpFile)
		defer ip_f.Close()
		if err != nil {
			return err
		}
		ip_buf := bufio.NewReader(ip_f)

		for {
			ip, err := ip_buf.ReadString(10)
			ip = strings.TrimSpace(ip)
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}

			for _, scheme := range vars.Schemes {
				handled_set := HandleCustomPorts(*vars.Host, ip)
				for _,item := range handled_set{
					task := Task{
						Uri:  fmt.Sprintf("%s://%s", scheme, item.IP),
						Host: item.Host,
					}
					// 生产者，不断地往taskChan channel发送数据，直到channel阻塞
					taskChan <- task
				}
			}
		}
	}else if taskType == "ip_hostfile" {
		host_f, err := os.Open(*vars.HostFile)
		defer host_f.Close()
		if err != nil {
			return err
		}
		host_buf := bufio.NewReader(host_f)
		for {
			host, err := host_buf.ReadString(10)
			host = strings.TrimSpace(host)
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}
			for _,ip := range ip_list{
				for _, scheme := range vars.Schemes {
					handled_set := HandleCustomPorts(host, ip)
					for _,item := range handled_set{
						task := Task{
							Uri:  fmt.Sprintf("%s://%s", scheme, item.IP),
							Host: item.Host,
						}
						// 生产者，不断地往taskChan channel发送数据，直到channel阻塞
						taskChan <- task
					}
				}
			}

		}
	}else if taskType == "ipfile_hostfile" {
		ip_f, err := os.Open(*vars.IpFile)
		defer ip_f.Close()
		if err != nil {
			return err
		}
		ip_buf := bufio.NewReader(ip_f)

		host_f, err := os.Open(*vars.HostFile)
		defer host_f.Close()
		if err != nil {
			return err
		}
		host_buf := bufio.NewReader(host_f)

		for {
			ip, err := ip_buf.ReadString(10)
			ip = strings.TrimSpace(ip)
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}

			for {
				host, err := host_buf.ReadString(10)
				host = strings.TrimSpace(host)
				if err != nil {
					if err == io.EOF {
						break
					}
					return err
				}

				for _, scheme := range vars.Schemes {
					handled_set := HandleCustomPorts(host, ip)
					for _,item := range handled_set{
						task := Task{
							Uri:  fmt.Sprintf("%s://%s", scheme, item.IP),
							Host: item.Host,
						}
						// 生产者，不断地往taskChan channel发送数据，直到channel阻塞
						taskChan <- task
					}
				}
			}
		}
	}

	close(taskChan)
	wg.Wait()

	return nil
}
