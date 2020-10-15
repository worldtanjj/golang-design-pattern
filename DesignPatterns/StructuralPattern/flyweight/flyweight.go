package flyweight

import (
	"fmt"
)

type CharacterStyleUnit struct {
	font     int
	size     int
	colorRGB int
}

type CharacterStyleUnitFactory struct{}

var characterStyleMap = make(map[string]*CharacterStyleUnit)

func (this CharacterStyleUnitFactory) Get(font, size, colorRGB int) *CharacterStyleUnit {
	var key = fmt.Sprintf("%d_%d_%d", font, size, colorRGB)
	val, ok := characterStyleMap[key]
	if ok {
		return val
	}

	var newStyle = &CharacterStyleUnit{font: font, size: size, colorRGB: colorRGB}
	characterStyleMap[key] = newStyle

	return newStyle
}

type Character struct {
	style *CharacterStyleUnit
	data  string
}

func NewCharacter(style *CharacterStyleUnit, data string) *Character {
	return &Character{
		style: style,
		data:  data,
	}
}
