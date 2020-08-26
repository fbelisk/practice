package main

import "sync"

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(b func()) {
	w.Add(1)
	go func() {
		b()
		w.Done()
	}()
}
