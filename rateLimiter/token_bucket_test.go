package rateLimiter

import (
	"fmt"
	"testing"
	"time"
)

func TestBucket_Put(t *testing.T) {
	type fields struct {
		closeCh  chan int
		Capacity uint64
		Tokens   uint64
		Interval time.Duration
		inc      uint64
	}
	type args struct {
		inc uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			fields: fields{
				Capacity: 1000,
				Tokens:   0,
				Interval: 0,
				inc:      0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bucket{
				closeCh:  tt.fields.closeCh,
				Capacity: tt.fields.Capacity,
				Tokens:   tt.fields.Tokens,
				Interval: tt.fields.Interval,
				Inc:      tt.fields.inc,
			}
			for i:= 0; i < 500; i++ {
				go b.Put(1)
			}
			//for i:= 0; i < 499; i++ {
			//	go b.Take(0)
			//}
			fmt.Printf("%+v", b)
		})
	}
}