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
)


type Task struct {
	Uri		string
	Host 	string
}

func getTitle(body string) string{
	re := regexp.MustCompile(`<title>([\s\S]*?)</title>`)
	match := re.FindStringSubmatch(body)
	if match != nil && len(match) > 1{
		return strings.TrimSpace(match[1])
	}else{
		return ""
	}
}

func goScan(taskChan chan Task, wg *sync.WaitGroup){
	defer wg.Done()
	for {
		select {
		case task, ok := <-taskChan:
			if !ok {
				return
			} else {
				vars.ProcessBar.Add(1)
				body := utils.GetHttpBody(task.Uri, task.Host)
				title := getTitle(body)
				var result models.Result
				result.Uri = task.Uri
				result.Host = task.Host
				result.Title = title
				resultStr, _ := json.Marshal(result)
				if len(title) > 0{
					elog.Notice(fmt.Sprintf("Uri: %s, Host: %s --> %s", task.Uri, task.Host, title))
					utils.WriteLine(string(resultStr), *vars.OutFile)
				}else{
					elog.Warn(fmt.Sprintf("Uri: %s, Host: %s No title found", task.Uri, task.Host))
				}
			}
		}
	}
}
