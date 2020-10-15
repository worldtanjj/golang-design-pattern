package obserser

import (
	"fmt"
)

//被观察者
type Subject interface {
	registerObserver(ob *Observer)
	removeObserver(ob *Observer)
	notifyObserver(msg string)
}

type ConcreteSubjectOne struct {
	observers []*Observer
}

func NewConcreteSubjectOne() *ConcreteSubjectOne {
	return &ConcreteSubjectOne{
		observers: make([]*Observer, 0),
	}
}

func (this *ConcreteSubjectOne) removeObserver(ob *Observer) {
	ret := make([]*Observer, 0, len(this.observers))
	for _, val := range this.observers {
		if val != ob {
			ret = append(ret, val)
		}
	}
	this.observers = ret
}

func (this *ConcreteSubjectOne) registerObserver(ob *Observer) {
	this.observers = append(this.observers, ob)
}

func (this *ConcreteSubjectOne) notifyObserver(msg string) {
	for _, ob := range this.observers {
		ob.Update(msg)
	}
}

//Observer 观察者
type Observer struct {
	name string
}

func NewObserver(name string) *Observer {
	return &Observer{name: name}
}

func (this *Observer) Update(msg string) {
	fmt.Printf("%s 收到通知:%s\n", this.name, msg)
}
