package logger

// Configure a logger, keep the timestamps at the front except make them green

import (
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warn *log.Logger
	Error   *log.Logger
	Success   *log.Logger
	Req   *log.Logger
)

var Nofodoso string = "hi"

type Colors struct {
	Red string
	Blue string
	Green string
	Yellow string
	Purple string
	Reset string
}

var Clr Colors = Colors{
		Red: "\033[1;31m",
		Blue: "\033[1;34m",
		Green: "\033[1;32m",
		Yellow: "\033[1;33m",
		Purple: "\033[1;35m",
		Reset: "\033[0m",
}

func init() {
	Info = log.New(os.Stdout, Clr.Blue + "INFO: \033[0m", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(os.Stdout, Clr.Yellow + "WARNING: \033[0m", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, Clr.Red + "ERROR: \033[0m", log.Ldate|log.Ltime|log.Lshortfile)
	Success = log.New(os.Stderr, Clr.Green + "SUCCESS: \033[0m", log.Ldate|log.Ltime|log.Lshortfile)
	Req = log.New(os.Stderr, Clr.Purple + "REQUEST: \033[0m", log.Ldate|log.Ltime|log.Lshortfile)
}

