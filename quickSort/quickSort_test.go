package quickSort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	node := QuickSort{
		Array:[]int{33, 12, 11, 24,29,19,47,78,99,71,47},
	}
	node.Sort(0, len(node.Array) - 1)
	fmt.Printf("slice=%v\n",node.Array)
}