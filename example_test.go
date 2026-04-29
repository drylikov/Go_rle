
package rle_test

import (
	"fmt"

	"github.com/drylikov/go_rle"
)

func Example() {
	b := rle.EncodeInt64([]int64{1, 1, 1, 1, 1, 1, 1, 1, 125, 1, 1, 1, 1, 1, 1, 1, 1})
	fmt.Printf("buf: %#x\n", b)
	fmt.Printf("len: %v\n", len(b))
	// Output:
	// buf: 0x0210fa01020210
	// len: 7
}
