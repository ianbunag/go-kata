package benchmark

import (
	"log"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func getFunctionName(function interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()), ".")

	return strs[len(strs)-1]
}

func Time(function func() interface{}) {
	start := time.Now()
	defer func() {
		log.Printf("%s, execution time %s\n", getFunctionName(function), time.Since(start))
	}()

	function()
}
