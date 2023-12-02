package loggify

import "runtime"

const (
    LEVEL_INFO = iota
    LEVEL_WARN
    LEVEL_ERROR
    LEVEL_DEBUG

    info_line = "[INFO] "
    warn_line = "[WARN] "
    error_line = "[ERROR] "
    debug_line = "[DEBUG] "
)

var Level = 0

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
    pc, _, line, ok := runtime.Caller(1)
    if !ok {
        println("Loggify error")
        return
    }
    name := runtime.FuncForPC(pc).Name()
    println(str, name, line)
}
