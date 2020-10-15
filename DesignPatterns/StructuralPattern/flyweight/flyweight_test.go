package flyweight

import (
	"fmt"
	"testing"
)

func TestFlyweight(t *testing.T) {
	chars := make([]*Character, 0)

	var styleFactory = &CharacterStyleUnitFactory{}
	chars = append(chars, NewCharacter(styleFactory.Get(11, 22, 33), "data1"))
	chars = append(chars, NewCharacter(styleFactory.Get(11, 22, 33), "data2"))
	chars = append(chars, NewCharacter(styleFactory.Get(11, 22, 33), "data3"))
	chars = append(chars, NewCharacter(styleFactory.Get(11, 22, 33), "data4"))

	var ref1 = chars[0].style
	var ref2 = chars[1].style
	var ref3 = chars[2].style
	var ref4 = chars[3].style
	fmt.Printf("ref of data's style: %x\n%x\n%x\n%x\n", &ref1, &ref2, &ref3, &ref4)
}
