package facade

import "testing"

var expect = "A module running\nB module running"

func TestFacadeAPI(t *testing.T) {
	var api = NewApi()
	ret := api.Test()

	if ret != expect {
		t.Fatalf("expect %s, return %s", expect, ret)
	}
}
