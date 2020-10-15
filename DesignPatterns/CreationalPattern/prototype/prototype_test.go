package prototype

import (
	"fmt"
	"testing"
)

var manager *PrototypeManager

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t
	return &tc
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() Cloneable {
	tc := *t
	return &tc
}

func init() {
	manager = NewPrototypeManager()
	t1 := &Type1{
		name: "type1",
	}
	manager.Set("t1", t1)
}
func TestClone(t *testing.T) {
	t1 := manager.Get("t1")
	t2 := t1.Clone()

	fmt.Printf("ref t1:%v", &t1)
	fmt.Printf("ref t2:%v", &t2)
	if t1 == t2 {
		t.Fatal("error! get clone not working")
	}
}

func TestCloneFromManager(t *testing.T) {
	c := manager.Get("t1").Clone()
	fmt.Printf("type of c: %T", c)

	t1, ok := c.(*Type1)
	if t1.name != "type1" {
		t.Fatal("error")
	}
	if !ok {
		t.Fatal("error")

	}
}
