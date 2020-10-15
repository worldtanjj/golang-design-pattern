package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	var slice = make(Demo1Slice, 0)
	slice = append(slice, "aa")
	slice = append(slice, "bb")
	slice = append(slice, "cc")

	var it Iterator = &Demo1Iterator{
		cursor:  0,
		arrlist: slice,
	}

	for ; it.haxNext(); it.next() {
		fmt.Println(it.currentItem())
	}

}
