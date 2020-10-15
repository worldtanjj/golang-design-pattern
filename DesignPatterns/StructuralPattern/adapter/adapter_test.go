package adapter

import "testing"

func TestAdapter(t *testing.T) {
	var old = &oldClass{}
	var adapter IStandard = NewOldClassAdapter(old)

	adapter.Request()
}
