package main

// #cgo LDFLAGS: -ldl
// #include <time.h>
// #include "hook.h"
import "C"
import (
	"os"
	"strconv"
)

var offsetSec int64
var offsetNSec int64

func init() {
	offsetSec = getEnvAsInt("FOOLPROC_OFFSET_SEC")
	offsetNSec = getEnvAsInt("FOOLPROC_OFFSET_NSEC")
}

func getEnvAsInt(key string) int64 {
	value := os.Getenv(key)
	if value == "" {
		return 0
	}
	valueAsInt, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err.Error())
	}
	return valueAsInt
}

//export on_clock_gettime
func on_clock_gettime(clk_id C.clockid_t, tp *C.struct_timespec) {
	tp.tv_sec = tp.tv_sec + C.__time_t(offsetSec)
	tp.tv_nsec = tp.tv_nsec + C.__syscall_slong_t(offsetNSec)
}

func main() {
}
