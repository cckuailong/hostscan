package utils

import (
	"fmt"
	"hostscan/elog"
	"io/ioutil"
	"strings"
	"os"
)

func Readlines(filepath string) []string{
	data, err := ioutil.ReadFile(filepath)
	if err != nil{
		elog.Error(fmt.Sprintf("Read File: %s [%s]", filepath, err))
		return nil
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	return lines
}

func WriteLine(line string, outpath string){
	f, err := os.OpenFile(outpath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(line+"\n"); err != nil {
		elog.Warn(fmt.Sprintf("Write uri[%s]: %s", line, err))
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}