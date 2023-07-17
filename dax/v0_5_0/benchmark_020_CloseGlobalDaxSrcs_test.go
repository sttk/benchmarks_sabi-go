package v0_5_0_test

import (
	"testing"

	sabi "github.com/sttk/benchmarks_sabi/dax/v0_5_0"
)

func Benchmark_CloseGlobalDaxSrcs_zeroDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
	}
	b.StopTimer()
}

func Benchmark_CloseGlobalDaxSrcs_oneDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.CloseGlobalDaxSrcs()
	}
	b.StopTimer()
}

func Benchmark_CloseGlobalDaxSrcs_freeDs(b *testing.B) {
	b.StopTimer()
	sabi.ResetGlobals()
	defer sabi.ResetGlobals()

	sabi.AddGlobalDaxSrc("cliargs", FooDaxSrc{})
	sabi.AddGlobalDaxSrc("database", FooDaxSrc{})
	sabi.AddGlobalDaxSrc("pubsub", FooDaxSrc{})
	sabi.AddGlobalDaxSrc("json", FooDaxSrc{})
	sabi.AddGlobalDaxSrc("env", FooDaxSrc{})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sabi.CloseGlobalDaxSrcs()
	}
	b.StopTimer()
}