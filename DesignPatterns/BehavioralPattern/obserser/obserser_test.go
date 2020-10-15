package obserser

import "testing"

func TestObserver(t *testing.T) {
	var sub Subject = NewConcreteSubjectOne()
	var ob1 = NewObserver("张三")
	var ob2 = NewObserver("李四")

	sub.registerObserver(ob1)
	sub.registerObserver(ob2)
	sub.registerObserver(ob2)
	sub.notifyObserver("通知了")

	t.Fail()
}
