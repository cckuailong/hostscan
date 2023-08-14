package utils

import (
	"crypto/tls"
	"fmt"
	"hostscan/vars"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetHttpBody(url, host string) string{
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout: time.Duration(*vars.Timeout) * time.Second,
	}

	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		//elog.Error(fmt.Sprintf("DoGet: %s [%s]", url, err))
		return ""
	}

	reqest.Host = host

	var ua string
	if *vars.IsRandUA == true {
		ua = RandUA()
	} else{
		ua = fmt.Sprintf("golang-hostscan/%v", vars.Version)
	}

	reqest.Header.Add("User-Agent", ua)

	response, err := client.Do(reqest)
	if response != nil{
		defer response.Body.Close()
	}
	
	if err != nil {
		//elog.Error(fmt.Sprintf("DoGet: %s [%s]", url, err))
		return ""
	}

	filter_status_codes := []int{}
	filters := strings.TrimSpace(*vars.FilterRespStatusCodes)
	if len(filters) > 0{
		for _,status_code := range strings.Split(filters, ","){
			filter_status_code, err := strconv.Atoi(strings.TrimSpace(status_code))
			if err != nil{
				continue
			}
			filter_status_codes = append(filter_status_codes, filter_status_code)
		}
		if !containsStatusCode(response.StatusCode, filter_status_codes){
			return ""
		}
	}

	bodyByte, _ := ioutil.ReadAll(response.Body)
	body := string(bodyByte)

	return body
}

func containsStatusCode(a int, l []int) bool {
	for _,item := range l{
		if a == item{
			return true
		}
	}

	return false
}
