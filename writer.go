package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	debugWriteModel   *WriteModel
	infoWriteModel    *WriteModel
	warningWriteModel *WriteModel
	errorWriteModel   *WriteModel
	fatalWriteModel   *WriteModel
)

type WriteModel struct {
	file        *os.File
	log_buffer  chan string
	exit_buffer chan bool
}

func (w *WriteModel) write(info interface{}) {
	msg := fmt.Sprintf("%v", info)
	w.exit_buffer <- true
	w.log_buffer <- msg
}

func (w *WriteModel) output() {
	go func(w *WriteModel) {
		do_logger := log.New(w.file, "\r\n" /*log.Ldate|*/, log.Ltime|log.Llongfile)
		var loginfo string
		for {
			loginfo = <-w.log_buffer
			do_logger.Print(loginfo)
			<-w.exit_buffer
		}
	}(w)

}
