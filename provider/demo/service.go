package demo

import "webframeworkV2.0/framework"

type DemoService struct {
	Service

	c framework.Container
}

func (s *DemoService) GetFoo() Foo {
	return Foo{
		Name: "i am foo",
	}
}

func NewDemoService(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)

	return &DemoService{c: c}, nil
}
