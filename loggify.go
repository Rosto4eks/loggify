package loggify

import (
	"runtime"
	"time"
)

const (
    LEVEL_INFO = iota
    LEVEL_WARN
    LEVEL_ERROR
    LEVEL_DEBUG
)

const (
	reset  = "\033[0;0m"
	red    = "\033[1;31m"
	green  = "\033[1;32m"
	yellow = "\033[1;33m"
	cyan   = "\033[1;36m"
	purple = "\033[1;35m"
)

const (
    info_line = green + "[INFO]  " + reset
    warn_line = yellow + "[WARN]  " + reset
    error_line = red + "[ERROR] " + reset
    debug_line = cyan + "[DEBUG] " + reset

    func_line = purple + "FUNC:  " + reset
    time_line = purple + "TIME:  " + reset
)


var level = 0

func SetLevel(lvl int) {
    if lvl < LEVEL_INFO || lvl > LEVEL_DEBUG {
        panic("incorrect level for loggify")
    }
    level = lvl
}

func INFO(str string) {
    log(info_line + str, LEVEL_INFO)
}

func WARN(str string) {
    log(warn_line + str, LEVEL_WARN)
}

func ERROR(str string) {
    log(error_line + str, LEVEL_ERROR)
}

func DEBUG(str string) {
    log(debug_line + str, LEVEL_DEBUG)
}

func log(str string, lvl int) {
    if lvl < level {
        return 
    }
    pc, _, line, ok := runtime.Caller(2)
    if !ok {
        println("Loggify error")
        return
    }
    name := runtime.FuncForPC(pc).Name()
    println(str)
    println(func_line, name, line)
    println(time_line, time.Now().Format("2 Jan 15:04:05"))
    println()
}
