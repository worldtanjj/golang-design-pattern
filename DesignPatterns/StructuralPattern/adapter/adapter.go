package adapter

import (
	"fmt"
)

type oldClass struct {
}

func (this oldClass) OldMethod() bool {
	fmt.Printf("oldmethod\n")
	return true
}

type IStandard interface {
	Request()
}

type OldClassAdapter struct {
	old *oldClass
}

func NewOldClassAdapter(old *oldClass) *OldClassAdapter {
	return &OldClassAdapter{
		old: old,
	}
}
func (this OldClassAdapter) Request() {
	r := this.old.OldMethod()
	fmt.Printf("旧接口返回:%v\n", r)

}
