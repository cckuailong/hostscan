package core

import (
	"encoding/json"
	"fmt"
	"hostscan/elog"
	"hostscan/models"
	"hostscan/utils"
	"hostscan/vars"
	"regexp"
	"strings"
	"sync"

	"github.com/schollz/progressbar/v3"
)

func getTitle(body string) string {
	re := regexp.MustCompile(`<title>([\s\S]*?)</title>`)
	match := re.FindStringSubmatch(body)
	if match != nil && len(match) > 1 {
		return strings.TrimSpace(match[1])
	} else {
		return ""
	}
}

func getTasks() [][2]string {
	tasks := [][2]string{}
	for _, ip := range vars.Ips {
		for _, scheme := range vars.Schemes {
			uri := fmt.Sprintf("%s://%s", scheme, ip)
			for _, host := range vars.Hosts {
				tasks = append(tasks, [2]string{uri, host})
			}
		}
	}

	return tasks
}

func goScan(taskChan chan [2]string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case task, ok := <-taskChan:
			if !ok {
				return
			} else {
				vars.ProcessBar.Add(1)
				uri := task[0]
				host := task[1]
				body := utils.GetHttpBody(uri, host)
				title := getTitle(body)
				var result models.Result
				result.Uri = uri
				result.Host = host
				result.Title = title
				resultStr, _ := json.Marshal(result)
				if len(title) > 0 {
					elog.Notice(fmt.Sprintf("Uri: %s, Host: %s --> %s", uri, host, title))
					utils.WriteLine(string(resultStr), *vars.OutFile)
				} else {
					elog.Warn(fmt.Sprintf("Uri: %s, Host: %s No title found", uri, host))
				}
			}
		}
	}
}

func Scan() {
	tasks := getTasks()
	wg := sync.WaitGroup{}
	totalTask := len(tasks)
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
	taskChan := make(chan [2]string, *vars.Thread*4)

	// 创建vars.ThreadNum个协程
	for i := 0; i < *vars.Thread; i++ {
		go goScan(taskChan, &wg)
		wg.Add(1)
	}

	// 生产者，不断地往taskChan channel发送数据，直到channel阻塞
	for _, task := range tasks {
		taskChan <- task
	}
	close(taskChan)
	wg.Wait()
}
