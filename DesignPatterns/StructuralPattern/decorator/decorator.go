package decorator

type Component interface {
	Calc() int
}

type ConcreteComponent struct {
	num int
}

func (this ConcreteComponent) Calc() int {
	return 0
}

type MutiDecorator struct {
	c   Component
	num int
}

func (this MutiDecorator) Calc() int {
	return this.c.Calc()
}

func WrapMuti(c Component, n int) Component {
	return &MutiDecorator{
		c:   c,
		num: n,
	}
}

type AddDecorator struct {
	Component
	num int
}

func (this AddDecorator) Calc() int {
	return this.Component.Calc() + this.num
}

func NewAddDecorator(c Component, n int) Component {
	return &AddDecorator{
		Component: c,
		num:       n,
	}
}

type MulDecorator struct {
	Component
	num int
}

func (this MulDecorator) Calc() int {
	return this.Component.Calc() * this.num
}

func NewMulDecorator(c Component, n int) Component {
	return &MulDecorator{
		Component: c,
		num:       n,
	}
}
