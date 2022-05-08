// Package framework
//  服务容器的实现
package framework

import (
	"errors"
	"sync"
)

// Container 抽象服务容器
type Container interface {
	// Bind 绑定一个服务提供者 如果关键字凭证已经存在 会进行替换操作 返回error
	Bind(provider ServiceProvider) error

	// IsBind 判断关键字凭证是否已经绑定服务提供者
	IsBind(key string) bool

	// Make 根据关键字凭证获取一个服务实例
	Make(key string) (interface{}, error)

	// MustMake 根据关键字凭证返回一个服务实例
	MustMake(key string) interface{}

	// MakeNew 根据关键字凭证获取一个服务实例
	MakeNew(key string, params []interface{}) (interface{}, error)
}

// WebContainer 服务容器的具体实现
type WebContainer struct {
	// 一种写法 提示WebContainer要强制实现Container接口
	Container

	// providers 存储注册的服务提供者 key为字符串凭证
	providers map[string]ServiceProvider

	// instances 存储具体的实例 key为字符串凭证
	instances map[string]interface{}

	// lock用于锁住对容器的变更操作
	lock sync.RWMutex
}

func NewWebContainer() Container {
	return &WebContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

// Bind 将服务容器与关键字做绑定
func (w *WebContainer) Bind(provider ServiceProvider) error {
	w.lock.Lock()
	defer w.lock.Unlock()

	key := provider.Name()

	w.providers[key] = provider

	if provider.IsDefer() == false {
		// 判断是否在绑定的时候进行实例化
		// todo 增加绑定的时候进行实例化的代码
	}

	return nil
}

func (w *WebContainer) IsBind(key string) bool {
	return w.findServiceProvider(key) != nil
}

// Make 提供获取服务实例
func (w *WebContainer) Make(key string) (interface{}, error) {
	return w.make(key, nil, false)
}

func (w *WebContainer) MustMake(key string) interface{} {
	serv, err := w.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return serv
}

func (w *WebContainer) findServiceProvider(key string) ServiceProvider {
	// todo 这里锁了两次
	//w.lock.Lock()
	//defer w.lock.Unlock()

	if sp, ok := w.providers[key]; ok {
		// 已经注册了服务的提供者
		return sp
	}

	return nil
}

func (w *WebContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	if err := sp.Boot(w); err != nil {
		return nil, err
	}

	method := sp.Register(w)
	ins, err := method(w)
	if err != nil {
		return nil, err
	}
	return ins, nil
}

// make 真正的实例化一个服务
func (w *WebContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	w.lock.Lock()
	defer w.lock.Unlock()

	// 查询是否已经注册过了这个服务提供者 如果没有注册 返回错误
	sp := w.findServiceProvider(key)
	if sp == nil {
		// 没有提供服务的注册者 报错
		return nil, errors.New(key + "have not register")
	}

	if forceNew {
		// todo 增加强制实例化的代码
		return nil, nil
	}

	if ins, ok := w.instances[key]; ok {
		// 说明容器中已经实例化过
		// 直接使用实例化过的对象即可
		return ins, nil
	}

	// 说明容器中没有实例化过 则进行一次实例化
	inst, err := w.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}

	w.instances[key] = inst

	return inst, nil
}

func (w *WebContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	// todo
	return nil, nil
}
