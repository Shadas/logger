package logger

import (
// "log"
)

func Log(level int, infos ...interface{}) {
	if level < indexString(loggerConfig.LogLevel, levelNames) {
		return
	}

	switch level {
	case FATAL:
		if indexString(levelNames[FATAL], loggerConfig.LogRange) != -1 {
			fatalWriteModel.write(infos)
		}
		fallthrough
	case ERROR:
		if indexString(levelNames[ERROR], loggerConfig.LogRange) != -1 {
			errorWriteModel.write(infos)
		}
		fallthrough
	case WARNING:
		if indexString(levelNames[WARNING], loggerConfig.LogRange) != -1 {
			warningWriteModel.write(infos)
		}
		fallthrough
	case INFO:
		if indexString(levelNames[INFO], loggerConfig.LogRange) != -1 {
			infoWriteModel.write(infos)
		}
		fallthrough
	case DEBUG:
		if indexString(levelNames[DEBUG], loggerConfig.LogRange) != -1 {
			debugWriteModel.write(infos)
		}
	default:
		debugWriteModel.write(infos)
	}

}
