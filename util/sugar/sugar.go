package sugar

import (
	"dsa/util/color"
	"dsa/util/perf"
	"testing"
	"time"
)

// TODO: There is a bug in sugar which shows the test as PASSED when a test crashed mid execution

var ColoerTime = color.Blue
var ColorMem = color.Cyan
var ColorPassed = color.Green
var ColorFailed = color.Red
var ColorSkipped = color.Yellow

// Shock - Place as defer call at the first line within the anonymous testing.Run() function
//
// Advise: Place a `println()` above the resting.Run() function
//
// Provides colored test results, runtime, memory usage
func Shock(t *testing.T, ttName string) {
	startT := time.Now()
	print(ColoerTime + "TEST TIME ")
	perf.TimeTracker(startT, ttName)
	print(ColorMem + "TEST MEM  ")
	perf.PrintMemUsage(perf.KB, ttName)
	if t.Skipped() {
		print(ColorSkipped + "SKIPPED ")
	} else if !t.Failed() {
		print(ColorPassed + "PASSED ")
	} else {
		print(ColorFailed + "FAILED ")
	}
}

// Lite - Place as defer call at the first line within the anonymous testing.Run() function
//
// Advise: Place a `println()` above the resting.Run() function
//
// Provides colored test results, runtime
func Lite(t *testing.T, ttName string) {
	startT := time.Now()
	print(ColoerTime + "TEST TIME ")
	perf.TimeTracker(startT, ttName)
	if t.Skipped() {
		print(ColorSkipped + "SKIPPED ")
	} else if !t.Failed() {
		print(ColorPassed + "PASSED ")
	} else {
		print(ColorFailed + "FAILED ")
	}
}

// Zero - Place as defer call at the first line within the anonymous testing.Run() function
//
// Advise: Place a `println()` above the resting.Run() function
//
// Provides colored test results
func Zero(t *testing.T) {
	if t.Skipped() {
		print(ColorSkipped + "SKIPPED ")
	} else if !t.Failed() {
		print(ColorPassed + "PASSED ")
	} else {
		print(ColorFailed + "FAILED ")
	}
}

// Skip - Executes t.Skip() and adds color
func Skip(t *testing.T) {
	print(ColorSkipped + "SKIPPED")
	t.Skip()
}
