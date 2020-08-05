package a

import (
	"fmt"
	_ "unsafe"
)

//go:linkname say struct/b.Say
func say() {
	fmt.Println("ab")
}
