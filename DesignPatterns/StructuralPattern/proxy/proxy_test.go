package proxy

import (
	"strings"
	"testing"
)

func TestProxy(t *testing.T) {
	var sub Subject = &Proxy{}

	res := sub.Do()

	if !strings.Contains(res, "pre") {
		t.Fail()
	}
}
