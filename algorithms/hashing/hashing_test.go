package hashing

import (
	"dsa/util/perf"
	"testing"
	"time"
)

func TestSha256Simple(t *testing.T) {
	want := "test-sha"
	got := Sha256("test")
	if got != want {
		t.Errorf("Result was incorrect, got: %s, want: %s.", got, want)
	}
}

func TestSha256Table(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"hashing test", args{"test"}, "test-sha"},
		{"hashing abc", args{"abc"}, "test-sha"},
		{"hashing 123", args{"123"}, "test-sha"},
		{"hashing äöü", args{"äöü"}, "test-sha"},
	}
	for _, tt := range tests {
		startT := time.Now()
		t.Run(tt.name, func(t *testing.T) {
			if got := Sha256(tt.args.input); got != tt.want {
				t.Errorf("Sha256() = %v, want %v", got, tt.want)
			}
		})
		perf.TimeTracker(startT, tt.name)
		perf.PrintMemUsage(perf.KB, tt.name)
	}
}
