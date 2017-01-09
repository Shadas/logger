package logger

import (
// "log"
)

func Log(level int, info interface{}) {
	if level < indexString(loggerConfig.LogLevel, levelNames) {
		return
	}

	switch level {
	case FATAL:
		if indexString(levelNames[FATAL], loggerConfig.LogRange) != -1 {
			fatalWriteModel.write(info)
		}
		fallthrough
	case ERROR:
		if indexString(levelNames[ERROR], loggerConfig.LogRange) != -1 {
			errorWriteModel.write(info)
		}
		fallthrough
	case WARNING:
		if indexString(levelNames[WARNING], loggerConfig.LogRange) != -1 {
			warningWriteModel.write(info)
		}
		fallthrough
	case INFO:
		if indexString(levelNames[INFO], loggerConfig.LogRange) != -1 {
			infoWriteModel.write(info)
		}
		fallthrough
	case DEBUG:
		if indexString(levelNames[DEBUG], loggerConfig.LogRange) != -1 {
			debugWriteModel.write(info)
		}
	default:
		debugWriteModel.write(info)
	}

}
