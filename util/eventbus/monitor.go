// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

package eventbus

// A Monitor monitors the execution of a goroutine, allowing the caller to
// block until it is complete. The zero value of m is valid and its Wait method
// returns immediately.
type Monitor struct {
	done <-chan struct{} // immutable after initialization
}

// Wait blocks until the goroutine monitored by m has finished executing.
// It is safe to call Wait repeatedly, and from multiple concurrent goroutines.
func (m Monitor) Wait() {
	if m.done == nil {
		return
	}
	<-m.done
}

// To executes f in a new [Monitor]. The caller is responsible for waiting for
// the goroutine to complete.
func Go(f func()) Monitor {
	done := make(chan struct{})
	go func() { defer close(done); f() }()
	return Monitor{done: done}
}
