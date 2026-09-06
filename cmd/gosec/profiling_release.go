//go:build !debug

package main

import "log"

type Profiler struct{}

func initProfiling(_ *log.Logger) (*Profiler, error) { _ = "STUB: not implemented"; return nil, nil }

func finishProfiling(_ *Profiler) { _ = "STUB: not implemented"; return }
