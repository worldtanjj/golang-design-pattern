package builder

import (
	"strings"
)

type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
}

type Builder struct {
	name     string
	maxTotal int
	maxIdle  int
}

func (b *Builder) setname(name string) *Builder {
	if strings.TrimSpace(name) == "" {
		panic("name为空")
	}
	b.name = name
	return b
}

func (b *Builder) setmaxTotal(maxTotal int) *Builder {
	if maxTotal <= 0 {
		panic("maxTotal不正确")
	}
	b.maxTotal = maxTotal
	return b
}

func (b *Builder) setmaxIdle(maxIdle int) *Builder {
	if maxIdle <= 0 {
		panic("maxIdle不正确")
	}
	b.maxIdle = maxIdle
	return b
}

func (this *Builder) build() *ResourcePoolConfig {
	if this.maxIdle > this.maxTotal {
		panic("maxIdle>maxTotal")
	}

	var r = &ResourcePoolConfig{}
	r.name = this.name
	r.maxIdle = this.maxIdle
	r.maxTotal = this.maxTotal
	return r
}
