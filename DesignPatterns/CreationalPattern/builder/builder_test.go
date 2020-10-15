package builder

import "testing"

func TestBuild(t *testing.T) {
	var builder = &Builder{}
	var r = builder.setname("掌声").setmaxTotal(5).setmaxIdle(4).build()
	t.Fatal(r == nil)
}
