package log

import (
	"fmt"
	"os"
	"path"
	"server/constant"
	"sync"
	"time"
)

var fileMutex struct {
	sync.Mutex
}

func checkFolder() error {
	var err error = nil
	if _, err = os.Stat(constant.LOG_PATH); err == nil {
		return nil
	}
	fileMutex.Lock()
	defer fileMutex.Unlock()
	err = os.MkdirAll(constant.LOG_PATH, constant.LOG_PERMISSION)
	return err
}

func Log(logType string, msg string, args ...interface{}) {
	err := checkFolder()
	if err != nil {
		return
	} else if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	filepath := path.Join(constant.LOG_PATH, logType)
	timeStamp := time.Now().UTC().Format("2006-01-02 15:04:05")
	text := fmt.Sprintf("[%s]%s", timeStamp, msg)

	fileMutex.Lock()
	defer fileMutex.Unlock()
	log, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, constant.LOG_PERMISSION)
	if err != nil {
		return
	}
	defer log.Close()
	fmt.Fprintln(log, text)
}
