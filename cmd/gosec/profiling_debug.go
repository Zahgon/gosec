//go:build debug

package main

import (
	"flag"
	"log"
	"os"
	"sync"
)

var (
	flagCPUProfile = flag.String("cpuprofile", "", "write cpu profile to file")
	flagMemProfile = flag.String("memprofile", "", "write memory profile to file")
)

type Profiler struct {
	cpuProfileFile *os.File
	logger         *log.Logger
	cleanupOnce    sync.Once
	cpuProfile     string
	memProfile     string
}

func NewProfiler(cpuProfile, memProfile string, logger *log.Logger) *Profiler {
	_ = "STUB: not implemented"
	return nil
}

func (p *Profiler) Start() error { _ = "STUB: not implemented"; return nil }

func (p *Profiler) Stop() { _ = "STUB: not implemented"; return }

func (p *Profiler) writeMemoryProfile() error { _ = "STUB: not implemented"; return nil }

func initProfiling(logger *log.Logger) (*Profiler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func finishProfiling(profiler *Profiler) { _ = "STUB: not implemented"; return }
