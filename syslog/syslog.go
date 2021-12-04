package syslog

import (
	"log"
	"os"
	"time"
)

var logLevel = 0

const (
	DBG_LEVEL = 1
	STD_LEVEL = 2
	WAN_LEVEL = 3
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

func init() {
	time.LoadLocation("Asia/Seoul")
	dbg = log.New(os.Stdout, string(colorWhite)+"[DBG] "+string(colorReset), log.LstdFlags)
	std = log.New(os.Stdout, string(colorCyan)+"[STD] "+string(colorReset), log.LstdFlags)
	wan = log.New(os.Stdout, string(colorPurple)+"[WAN] "+string(colorReset), log.LstdFlags)
	err = log.New(os.Stdout, string(colorRed)+"[ERR] "+string(colorReset), log.LstdFlags)

	logLevel = DBG_LEVEL
}

func DBG(format string, v ...interface{}) {
	if logLevel > DBG_LEVEL {
		return
	}
	dbg.Printf(format, v...)
}

func ST(format string, v ...interface{}) {
	if logLevel > STD_LEVEL {
		return
	}
	std.Printf(format, v...)
}

func WARN(format string, v ...interface{}) {
	if logLevel > WAN_LEVEL {
		return
	}
	wan.Printf(format, v...)
}
func ERR(format string, v ...interface{}) {
	if logLevel > ERR_LEVEL {
		return
	}
	err.Printf(format, v...)
}

/*
func Log(msg string) {
	now := time.Now().Format("2006-01-02 15:04:05.000")
	fmt.Printf("[%s] %s\n", now, msg)
}
*/
