package demo

// DemoServiceProvider 服务的提供方
type DemoServiceProvider struct {
}

//
func (sp *DemoServiceProvider) Name() string {
	return Key
}
