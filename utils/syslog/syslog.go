package syslog

import (
	"log"
	"os"
	"time"
)

var logLevel int
var logLevelStr map[int]string

const (
	DBG_LEVEL = 1
	STD_LEVEL = 2
	WAR_LEVEL = 3
	ERR_LEVEL = 4
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

var (
	dbg *log.Logger
	std *log.Logger
	wan *log.Logger
	err *log.Logger
)

/* Flag option (https://pkg.go.dev/log#pkg-examples)
const (
   Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
   Ltime                         // the time in the local time zone: 01:23:23
   Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
   Llongfile                     // full file name and line number: /a/b/c/d.go:23
   Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
   LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
   Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
   LstdFlags     = Ldate | Ltime // initial values for the standard logger
)
*/

func init() {
	time.LoadLocation("Asia/Seoul")
	dbg = log.New(os.Stdout, string(colorWhite)+"[DBG] "+string(colorReset), log.LstdFlags)
	std = log.New(os.Stdout, string(colorCyan)+"[STD] "+string(colorReset), log.LstdFlags)
	wan = log.New(os.Stdout, string(colorPurple)+"[WAR] "+string(colorReset), log.LstdFlags)
	err = log.New(os.Stdout, string(colorRed)+"[ERR] "+string(colorReset), log.LstdFlags)

	logLevel = STD_LEVEL

	logLevelStr = make(map[int]string, 4)
	logLevelStr[DBG_LEVEL] = "DBG_LEVEL"
	logLevelStr[STD_LEVEL] = "STD_LEVEL"
	logLevelStr[WAR_LEVEL] = "WAR_LEVEL"
	logLevelStr[ERR_LEVEL] = "ERR_LEVEL"
}

func SetLogLevel(level int) {
	if level < DBG_LEVEL || level > ERR_LEVEL {
		WAR("Invalid log level: %d", level)
		return
	}
	STD("Change log level: [%s] -> [%s]", logLevelStr[logLevel], logLevelStr[level])
	logLevel = level
}

func DBG(format string, v ...interface{}) {
	if logLevel > DBG_LEVEL {
		return
	}
	dbg.Printf(format, v...)
}

func DBGLN(v ...interface{}) {
	if logLevel > DBG_LEVEL {
		return
	}
	dbg.Println(v...)
}

func STD(format string, v ...interface{}) {
	if logLevel > STD_LEVEL {
		return
	}
	std.Printf(format, v...)
}

func STDLN(v ...interface{}) {
	if logLevel > STD_LEVEL {
		return
	}
	std.Println(v...)
}

func WAR(format string, v ...interface{}) {
	if logLevel > WAR_LEVEL {
		return
	}
	wan.Printf(format, v...)
}

func WARLN(v ...interface{}) {
	if logLevel > WAR_LEVEL {
		return
	}
	wan.Println(v...)
}

func ERR(format string, v ...interface{}) {
	if logLevel > ERR_LEVEL {
		return
	}
	err.Printf(format, v...)
}

func ERRLN(v ...interface{}) {
	if logLevel > ERR_LEVEL {
		return
	}
	err.Println(v...)
}
