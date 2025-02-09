package perf

import (
	"dsa/util/color"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

var ColorTime = color.Blue
var ColorMem = color.Cyan

const (
	Byte uint64 = 1
	KB   uint64 = 1000
	MB   uint64 = 1000 * KB
	GB   uint64 = 1000 * MB
	KiB  uint64 = 1024
	MiB  uint64 = 1024 * KiB
	GiB  uint64 = 1024 * MiB
)

func nameOfByte(i uint64) string {
	switch i {
	case Byte:
		return "Byte"
	case KB:
		return "KB"
	case MB:
		return "MB"
	case GB:
		return "GB"
	case KiB:
		return "KiB"
	case MiB:
		return "MiB"
	case GiB:
		return "GiB"
	default:
		return "divMag: " + strconv.FormatUint(i, 10)
	}
}

func MeasurePerformance() {
	defer TimeTracker(time.Now(), "Main")
	defer PrintMemUsage(KB, "Main")

}

// PrintMemUsage usage: defer aocperf.PrintMemUsage(aocperf.KB, "Main")
// use name as identifier
// other memory magnitudes are available as constants in this module (e.g. aocperf.Byte, aocperf.MiB)
// Other magnitudes can also be specified -> TotalBytes / memMagnitude -> %g divMag: "memMagnitude"
func PrintMemUsage(memMagnitude uint64, name string) {
	n := nameOfByte(memMagnitude)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	out := "%s[ %s ] " +
		"Alloc = %g " + n +
		" | TotalAlloc = %g " + n +
		" | Sys = %g " + n +
		" | NumGC = %v\n"
	fmt.Printf(
		out,
		ColorMem,
		name,
		float64(m.Alloc)/float64(memMagnitude),
		float64(m.TotalAlloc)/float64(memMagnitude),
		float64(m.Sys)/float64(memMagnitude),
		m.NumGC,
	)
}

// TimeTracker use Time.Now() as start time. Name is the id later printed to std_out
// usage: defer aocperf.TimeTracker(time.Now(), "Main")
func TimeTracker(start time.Time, name string) {
	elapsed := time.Since(start)
	output := fmt.Sprintf("%s[ %s ] took %s\n", ColorTime, name, elapsed)
	print(output)
}
