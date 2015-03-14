package golibs

import (
	"testing"
	"fmt"
)

type MyType []int
func (a MyType) Len() int { return len(a) }
func (a MyType) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func TestShuffle(t *testing.T) {
	slice := []int{0,1,2,3,4,5,6,7,8,9}
	fmt.Printf("%v\n", slice)
	Shuffle(MyType(slice))
	fmt.Printf("%v", slice)
}

