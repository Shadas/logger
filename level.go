package logger

type Level int

const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
	FATAL
)

var levelNames = []string{
	"DEBUG",
	"INFO",
	"WARNING",
	"ERROR",
	"FATAL",
}

func indexString(s string, a []string) int {
	for index, as := range a {
		if s == as {
			return index
		}
	}
	return -1
}
