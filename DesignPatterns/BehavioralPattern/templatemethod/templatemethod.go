package templatemethod

import (
	"fmt"
)

type IPerson interface {
	SetName(name string)
	BeforeOut()
	Out()
}

type Person struct {
	name          string
	SpecifyAction IPerson
}

func (this *Person) SetName(name string) {
	this.name = name
}

func (this *Person) BeforeOut() {
	if this.SpecifyAction == nil {
		return
	}
	this.SpecifyAction.BeforeOut()
}

func (this *Person) Out() {
	this.BeforeOut()
	fmt.Println("出门")
}

type Boy struct {
	*Person
}

func (this *Boy) BeforeOut() {
	fmt.Println("睡一觉")
}

type Girl struct {
	*Person
}

func (this *Girl) BeforeOut() {
	fmt.Println("化妆1小时")
}
