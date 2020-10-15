package strategy

import (
	"fmt"
)

type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

func NewPayment(name, cardNo string, money float64, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardNo: cardNo,
			Money:  money,
		},
		strategy: strategy,
	}
}

type PaymentContext struct {
	Name   string
	CardNo string
	Money  float64
}

type PaymentStrategy interface {
	Pay(*PaymentContext)
}

func (this *Payment) Pay() {
	this.strategy.Pay(this.context)
}

type CashPay struct{}

func (this *CashPay) Pay(ctx *PaymentContext) {
	fmt.Printf("%v payed %.2f use cash\n", ctx.Name, ctx.Money)
}

type CardPay struct{}

func (this *CardPay) Pay(ctx *PaymentContext) {
	fmt.Printf("%v payed %.2f use card, cardno is: %v\n", ctx.Name, ctx.Money, ctx.CardNo)
}
