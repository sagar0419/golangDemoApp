package api

import (
	"fmt"
	"net/http"
	"runtime"
)

func Metrics(w http.ResponseWriter, r *http.Request) {
	cpuNum := runtime.NumCPU()
	memStats := new(runtime.MemStats)
	runtime.ReadMemStats(memStats)

	fmt.Fprintln(w, "Number of CPU:", cpuNum)
	fmt.Fprintln(w, "Memory consumption:", memStats.Alloc)
}
