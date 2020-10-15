package facade

import "fmt"

type API interface {
	Test() string
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (this apiImpl) Test() string {
	aRet := this.a.MethodA()
	bRet := this.b.MethodB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

func NewApi() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

type AModuleAPI interface {
	MethodA() string
}

type aModuleImpl struct{}

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

func (this aModuleImpl) MethodA() string {
	return "A module running"
}

type BModuleAPI interface {
	MethodB() string
}

type bModuleImpl struct{}

func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

func (this bModuleImpl) MethodB() string {
	return "B module running"
}
