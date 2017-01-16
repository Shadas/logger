package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	debugWriteModel   *WriteModel
	infoWriteModel    *WriteModel
	warningWriteModel *WriteModel
	errorWriteModel   *WriteModel
	fatalWriteModel   *WriteModel
)

type WriteModel struct {
	level       string
	file        *os.File
	log_buffer  chan string
	exit_buffer chan bool
}

func (w *WriteModel) write(infos ...interface{}) {
	var msg string
	for _, info := range infos {
		msg += fmt.Sprintf("%v", info)
	}

	w.exit_buffer <- true
	w.log_buffer <- msg
}

func (w *WriteModel) output() {
	go func(w *WriteModel) {
		if loggerConfig.SeparateFileByDate {
			now := time.Now().Format("2006-01-02")
			if !strings.Contains(w.file.Name(), now) {
				var err error
				w.file.Close()
				logfilename := loggerConfig.LogPath + "/" + now + "_" + w.level + ".log"
				w.file, err = os.OpenFile(logfilename, os.O_RDWR|os.O_CREATE, 0777)
				if err != nil {
					log.Panicln(err.Error())
				}
			}
		}
		do_logger := log.New(w.file, "\r\n" /*log.Ldate|*/, log.Ltime|log.Llongfile)
		var loginfo string
		for {
			loginfo = <-w.log_buffer
			do_logger.Print(loginfo)
			<-w.exit_buffer
		}
	}(w)

}
