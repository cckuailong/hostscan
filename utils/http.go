package utils

import (
	"crypto/tls"
	"fmt"
	"hostscan/vars"
	"io/ioutil"
	"net/http"
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
	bodyByte, _ := ioutil.ReadAll(response.Body)
	body := string(bodyByte)

	return body
}
