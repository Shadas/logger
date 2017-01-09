package logger

import (
	"log"
	"os"
	"time"
)

func InitLogger(config LoggerConfig) {
	loggerConfig = &config

	debugWriteModel = &WriteModel{}
	infoWriteModel = &WriteModel{}
	warningWriteModel = &WriteModel{}
	errorWriteModel = &WriteModel{}
	fatalWriteModel = &WriteModel{}

	logbuffer := loggerConfig.LogBuffer

	for _, level := range loggerConfig.LogRange {
		if indexString(level, levelNames) != -1 {
			file, err := os.OpenFile(loggerConfig.LogPath+"/"+level+".log", os.O_RDWR|os.O_CREATE, 0777)
			if err != nil {
				log.Panicln(err.Error())
			}

			switch level {
			case levelNames[DEBUG]:
				debugWriteModel.file = file
				debugWriteModel.log_buffer = make(chan string, logbuffer)
				debugWriteModel.exit_buffer = make(chan bool, logbuffer)
				debugWriteModel.output()
			case levelNames[INFO]:
				infoWriteModel.file = file
				infoWriteModel.log_buffer = make(chan string, logbuffer)
				infoWriteModel.exit_buffer = make(chan bool, logbuffer)
				infoWriteModel.output()
			case levelNames[WARNING]:
				warningWriteModel.file = file
				warningWriteModel.log_buffer = make(chan string, logbuffer)
				warningWriteModel.exit_buffer = make(chan bool, logbuffer)
				warningWriteModel.output()
			case levelNames[ERROR]:
				errorWriteModel.file = file
				errorWriteModel.log_buffer = make(chan string, logbuffer)
				errorWriteModel.exit_buffer = make(chan bool, logbuffer)
				errorWriteModel.output()
			case levelNames[FATAL]:
				fatalWriteModel.file = file
				fatalWriteModel.log_buffer = make(chan string, logbuffer)
				fatalWriteModel.exit_buffer = make(chan bool, logbuffer)
				fatalWriteModel.output()
			}
		}
	}

}

func CloseLogger() {
	if debugWriteModel != nil && debugWriteModel.file != nil {
		for {
			time.Sleep(1 * time.Microsecond)
			if len(debugWriteModel.exit_buffer) == 0 {
				break
			}
		}
		debugWriteModel.file.Close()
	}
	if infoWriteModel != nil && infoWriteModel.file != nil {
		for {
			time.Sleep(1 * time.Microsecond)
			if len(debugWriteModel.exit_buffer) == 0 {
				break
			}
		}
		infoWriteModel.file.Close()
	}
	if warningWriteModel != nil && warningWriteModel.file != nil {
		for {
			time.Sleep(1 * time.Microsecond)
			if len(debugWriteModel.exit_buffer) == 0 {
				break
			}
		}
		warningWriteModel.file.Close()
	}
	if errorWriteModel != nil && errorWriteModel.file != nil {
		for {
			time.Sleep(1 * time.Microsecond)
			if len(debugWriteModel.exit_buffer) == 0 {
				break
			}
		}
		errorWriteModel.file.Close()
	}
	if fatalWriteModel != nil && fatalWriteModel.file != nil {
		for {
			time.Sleep(1 * time.Microsecond)
			if len(debugWriteModel.exit_buffer) == 0 {
				break
			}
		}
		fatalWriteModel.file.Close()
	}
}
