package state

type State int

const (
	SMALL State = 0
	SUPER State = 1
	FIRE  State = 2
	CAPE  State = 3
)

type IMario interface {
	obtainMushRoom(state *MarioStateMachine)
	obtainCape(state *MarioStateMachine)
	obtainFireFlower(state *MarioStateMachine)
	meetMonster(state *MarioStateMachine)
}

type MarioStateMachine struct {
	score        int
	currentState IMario
}

func NewMarioStateMachine() *MarioStateMachine {
	return &MarioStateMachine{
		score:        0,
		currentState: &SmallMario{},
	}
}

func (this *MarioStateMachine) obtainMushRoom() {
	this.currentState.obtainMushRoom(this)
}
func (this *MarioStateMachine) obtainCape() {
	this.currentState.obtainCape(this)
}
func (this *MarioStateMachine) obtainFireFlower() {
	this.currentState.obtainFireFlower(this)

}
func (this *MarioStateMachine) meetMonster() {
	this.currentState.meetMonster(this)
}

type SmallMario struct{}

func (this *SmallMario) obtainMushRoom(state *MarioStateMachine) {
	state.score += 100
	state.currentState = &SuperMario{}
}
func (this *SmallMario) obtainCape(state *MarioStateMachine) {
	state.score += 200
	state.currentState = &CapeMario{}
}
func (this *SmallMario) obtainFireFlower(state *MarioStateMachine) {
	state.score += 300
	state.currentState = &FireMario{}
}
func (this *SmallMario) meetMonster(state *MarioStateMachine) {
	panic("game over")
}

type SuperMario struct{}

func (this *SuperMario) obtainMushRoom(state *MarioStateMachine) {
	state.score += 100
}
func (this *SuperMario) obtainCape(state *MarioStateMachine) {
	state.score += 200
	state.currentState = &CapeMario{}
}
func (this *SuperMario) obtainFireFlower(state *MarioStateMachine) {
	state.score += 300
	state.currentState = &FireMario{}
}
func (this *SuperMario) meetMonster(state *MarioStateMachine) {
	state.score -= 100
	state.currentState = &SmallMario{}
}

type FireMario struct{}

func (this *FireMario) obtainMushRoom(state *MarioStateMachine) {
	state.score += 100
}
func (this *FireMario) obtainCape(state *MarioStateMachine) {
	state.score += 200
}
func (this *FireMario) obtainFireFlower(state *MarioStateMachine) {
	state.score += 300
}
func (this *FireMario) meetMonster(state *MarioStateMachine) {
	state.score -= 300
	state.currentState = &SmallMario{}
}

type CapeMario struct{}

func (this *CapeMario) obtainMushRoom(state *MarioStateMachine) {
	state.score += 100
}
func (this *CapeMario) obtainCape(state *MarioStateMachine) {
	state.score += 200
}
func (this *CapeMario) obtainFireFlower(state *MarioStateMachine) {
	state.score += 300
}
func (this *CapeMario) meetMonster(state *MarioStateMachine) {
	state.score -= 200
	state.currentState = &SmallMario{}
}
