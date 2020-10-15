package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	var c = &ConcreteComponent{}
	var addc = NewAddDecorator(c, 2)
	var mutic = NewMulDecorator(addc, 5)

	var result = mutic.Calc()

	fmt.Printf("the result:%d\n", result)
}
