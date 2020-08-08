package rateLimiter

import (
	"sync/atomic"
	"time"
)

type Bucket struct {
	closeCh  chan int
	Capacity uint64
	Tokens   uint64
	Interval time.Duration
	inc      uint64
}

func (b *Bucket) Start() {
	ticker := time.NewTicker(b.Interval)
	defer ticker.Stop()

	for _ = range ticker.C {
		select {
		case <-b.closeCh:
			return
		default:
			b.Put(b.inc)
		}
	}
}

func (b *Bucket) Put(count uint64) {
	for {
		t := atomic.LoadUint64(&b.Tokens)
		if t == b.Capacity {
			return
		}
		if !atomic.CompareAndSwapUint64(&b.Tokens, t, t+count) {
			continue
		}
		return
	}
}

func (b *Bucket) Take(count uint64) bool{
	for {
		t := atomic.LoadUint64(&b.Tokens)
		if t == 0 {
			return false
		}
		if !atomic.CompareAndSwapUint64(&b.Tokens, t, t-count) {
			continue
		}
		return true
	}
}
