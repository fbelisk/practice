package slideWindows

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	maxs := MaxSlidingWindow1([]int{1,3,-1,-3,5,3,6,7}, 3)
	fmt.Printf("%v", maxs)
}
