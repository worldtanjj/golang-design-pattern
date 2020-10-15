package memento

type Memento interface{}

type Game struct {
	hp, mp int
}
type gameMemento struct {
	hp, mp int
}

func (this *Game) Play(h, m int) {
	this.hp += h
	this.mp += m
}

func (this *Game) Save() *gameMemento {
	return &gameMemento{
		hp: this.hp,
		mp: this.mp,
	}
}
func (this *Game) Load(m Memento) *Game {
	gm := m.(*gameMemento)
	this.hp = gm.hp
	this.mp = gm.mp
	return this
}
