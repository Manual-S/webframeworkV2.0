package demo

import (
	"fmt"

	"webframeworkV2.0/framework"
)

// DemoServiceProvider 服务的提供方
type DemoServiceProvider struct {
}

//
func (sp *DemoServiceProvider) Name() string {
	return Key
}

// Register 是注册初始化服务实例的方法
func (sp *DemoServiceProvider) Register(c framework.Container) framework.NewInstance {
	return nil
}

// IsDefer 表示是否延迟实例化 这里设置为true 表示延迟初始化
// 将这个服务的实例化延迟到第一次make的时候
func (sp *DemoServiceProvider) IsDefer() bool {
	return true
}

func (sp *DemoServiceProvider) Params(c framework.Container) []interface{} {
	return nil
}

func (sp *DemoServiceProvider) Boot(c framework.Container) error {
	fmt.Println("demo service boot")
	return nil
}
