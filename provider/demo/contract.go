package demo

const Key = "web:demo"

// 服务的接口
type Service interface {
	GetFoo() Foo
}

type Foo struct {
	Name string
}
