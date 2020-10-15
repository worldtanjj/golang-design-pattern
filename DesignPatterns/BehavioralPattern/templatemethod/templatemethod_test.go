package templatemethod

import "testing"

func TestTemplate(t *testing.T) {
	var p *Person = &Person{}
	p.SpecifyAction = &Boy{}
	p.SetName("小明")
	p.Out()

	p.SpecifyAction = &Girl{}
	p.SetName("小红")
	p.Out()
}
