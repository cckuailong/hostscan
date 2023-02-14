package utils

import (
	"bytes"
	"fmt"
	"hostscan/elog"
	"os"
)

func LineCounter(filepath string) (int, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		switch c, err := f.Read(buf[:]); true {
		case c < 0:
			return 0, err
		case c == 0: // EOF
			return count, nil
		case c > 0:
			count += bytes.Count(buf[:c], lineSep)
		}
	}

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