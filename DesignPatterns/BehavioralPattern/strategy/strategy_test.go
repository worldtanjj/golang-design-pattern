package strategy

import "testing"

func TestStrategy(t *testing.T) {
	var payment = NewPayment("aa", "", 1.22222, &CashPay{})
	payment.Pay()

	var payment1 = NewPayment("小子", "622", 100.001, &CardPay{})
	payment1.Pay()
}
