package framework

// NewInstance 定义了如何创建一个新实例
type NewInstance func(...interface{}) (interface{}, error)

// ServiceProvider 定义了一个服务提供者需要实现的接口
type ServiceProvider interface {
	// Name 代表了这个服务提供者的凭证
	Name() string

	// Register 在服务容器中注册了一个实例化服务的方法 是否在注册的时候
	// 就实例化这个服务 需要参考IsDefer接口
	Register(Container) NewInstance

	Params(Container) []interface{}

	// IsDefer 决定是否在注册的时候示例化这个服务
	IsDefer() bool

	// Boot 在调用实例化服务的时候会调用
	Boot(Container) error
}
