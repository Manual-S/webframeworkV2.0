package demo

const Key = "web:demo"

type Service interface {
	GetFoo() Foo
}

type Foo struct {
	Name string
}
