package facade

import "fmt"

//外观模式
type API interface {
	Test() string
}

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

type AModuleAPI interface {
	TestA() string
}

type BModuleAPI interface {
	TestB() string
}

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

type aModuleImpl struct{}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

type bModuleImpl struct{}

func (*bModuleImpl) TestB() string {
	return "B module running"
}