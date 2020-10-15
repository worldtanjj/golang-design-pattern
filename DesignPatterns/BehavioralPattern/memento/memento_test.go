package memento

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	var game = &Game{hp: 10, mp: 10}

	var m = game.Save()
	game.Play(1, 2)
	fmt.Println(game)

	game.Play(4, 4)
	fmt.Println(game)

	game.Load(m)

	fmt.Println(game)
}
