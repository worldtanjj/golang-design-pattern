package simplefactory

import "testing"

func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("tom")
	if s != "Hi, tom" {
		t.Fatal("测试失败")
	}
}

func TestNil(t *testing.T) {
	api := NewAPI(3)
	if api != nil {
		t.Fatal("nil")
	}
}
