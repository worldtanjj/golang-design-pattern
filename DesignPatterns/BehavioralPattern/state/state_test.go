package state

import (
	"fmt"
	"testing"
)

func TestState(t *testing.T) {
	var mario = NewMarioStateMachine()
	mario.obtainMushRoom()
	mario.obtainFireFlower()
	mario.obtainCape()
	mario.meetMonster()
	//mario.meetMonster()

	fmt.Printf("积分:%v,当前状态 %T\n", mario.score, mario.currentState)
}
