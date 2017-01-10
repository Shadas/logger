package logger

type LoggerConfig struct {
	LogPath            string   `json:"log_path"`
	LogLevel           string   `json:"log_level"`
	LogRange           []string `json:"log_range"`
	LogBuffer          int      `json:"log_buffer"`
	SeparateFileByDate bool     `json:"separate_file_by_date"`
}

var loggerConfig *LoggerConfig
