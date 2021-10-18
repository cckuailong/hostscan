package utils

import (
	"crypto/tls"
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
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
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
